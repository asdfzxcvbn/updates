package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

type ITunesResult struct {
	Name    string `json:"trackName"`
	Version string `json:"version"`
}

type ITunesResponse struct {
	Results []ITunesResult `json:"results"`
}

var (
	AppIdRegex = regexp.MustCompile(`(?:id)(\d{9,10})`)

	ErrNoAppId   = errors.New("no app id found")
	ErrNoResults = errors.New("no results found")
)

func AppIdFromLink(link string) (string, error) {
	full := AppIdRegex.FindStringSubmatch(link)
	if len(full) < 2 {
		return "", ErrNoAppId
	}
	return full[1], nil
}

func getLatestInfo(appid string) (*ITunesResult, error) {
	url := fmt.Sprintf("https://itunes.apple.com/lookup?requestId=%s&limit=1&id=%s", rand.Text(), appid)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var itunesResp ITunesResponse
	if err = json.NewDecoder(resp.Body).Decode(&itunesResp); err != nil {
		return nil, err
	}

	if len(itunesResp.Results) == 0 {
		return nil, ErrNoResults
	}

	return &itunesResp.Results[0], nil
}
