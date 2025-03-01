# updates
finally, a well-written appstore update notifier (in go) !!

## setup instructions
i'll assume you'll use a prebuilt binary, since you already know what to do if you want to build this yourself.

1. download the appropriate binary for your system from the [latest release](https://github.com/asdfzxcvbn/updates/releases/latest)
2. make a file named `config.toml` (or not, you can specify the config path yourself using `-config`) and base it off the [template](https://github.com/asdfzxcvbn/updates/blob/main/config.toml)
3. start checking for updates!

```
$ ./updates -config ./config.toml
```

## features/notes
* `updates` automatically detects changes to your `config.toml` and refreshes it at runtime. you **don't** have to restart the program to add links to more apps!

* `updates` validates the bot token and chat id provided.

* you can provide your own message template! telegram parses the message using html, so you can use html tags! you can decide to not specify some values, if you don't want them, like the appstore link.