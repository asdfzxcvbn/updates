package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/asdfzxcvbn/updates/db"
)

func main() {
	for {
		cancelled := false

		for _, link := range config.Apps {
			appid, err := AppIdFromLink(link)
			if err != nil {
				log.Fatalf("error checking %s: %v", link, err)
			}

			log.Printf("checking %s for updates in 5 seconds...", link)
			time.Sleep(5 * time.Second)

			//
			// step 1: get latest app version and compare to current known version
			//

			latestInfo, err := getLatestInfo(appid)
			if err != nil {
				log.Printf("error fetching %s: %v", link, err)
				log.Fatalln("this likely indicates the app no longer exists on the app store. it may also be the iTunes api's fault, so you should check if the app link is still valid.")
			}

			currentVersion, err := dbQueries.GetCurrentVersion(dbCtx, appid)
			if errors.Is(err, sql.ErrNoRows) { // this means this app was just added and we just fetched the initial version.
				if err = dbQueries.InsertVersion(dbCtx, db.InsertVersionParams{ID: appid, Version: latestInfo.Version}); err != nil {
					log.Fatalf("error inserting newest version for %s: %v", link, err)
				}

				log.Printf("inserted version for new app %s: %s\n\n", link, latestInfo.Version)
				continue
			} else if err != nil {
				log.Fatalf("error getting current version for %s: %v", link, err)
			}

			if currentVersion == latestInfo.Version {
				log.Printf("no update detected, keeping version %s\n\n", currentVersion)
				continue
			}

			//
			// step 2: update found, notify!
			//

			log.Printf("found an update!! notifying about %s -> %s ...\n", currentVersion, latestInfo.Version)

			if err = dbQueries.UpdateVersion(dbCtx, db.UpdateVersionParams{ID: appid, Version: latestInfo.Version}); err != nil {
				log.Fatalf("couldn't update db: %v", err)
			}

			replacer := strings.NewReplacer(
				"{app_name}", latestInfo.Name,
				"{old_version}", currentVersion,
				"{new_version}", latestInfo.Version,
				"{appstore_link}", link,
			)
			if err = messenger.SendMessage(replacer.Replace(config.Template)); err != nil {
				log.Fatalf("couldn't notify: %v", err)
			}

			log.Print("done!\n\n")

			// this should be at the end to make sure it's checked even if the config file is changed while the last app is being checked.
			// this will prevent us from having to wait 5 minutes before we start checking new apps from the config.
			if err = uCtx.Err(); err != nil {
				cancelled = true
				uCtx, uCtxCancel = context.WithCancel(context.Background())
				break
			}
		}

		if cancelled {
			log.Print("config file was updated. checking all apps again!\n\n")
			continue
		}

		log.Print("done checking all apps. restarting in 5 minutes!\n\n")
		time.Sleep(5 * time.Minute)
	}
}
