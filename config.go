package main

import (
	"context"
	"flag"

	"github.com/asdfzxcvbn/tgmessenger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	BotToken string   `mapstructure:"bot_token"`
	DbPath   string   `mapstructure:"db_path"`
	ChatId   string   `mapstructure:"chat_id"`
	TopicId  int64    `mapstructure:"topic_id"`
	Template string   `mapstructure:"template"`
	Apps     []string `mapstructure:"apps"`
}

var (
	uViper = viper.New()

	uCtx       context.Context
	uCtxCancel context.CancelFunc

	config     Config
	configPath string
)

// first time config load
func init() {
	flag.StringVar(&configPath, "config", "./config.toml", "path to config file")
	flag.Parse()

	uCtx, uCtxCancel = context.WithCancel(context.Background())

	uViper.SetConfigFile(configPath)
	reloadConfig(false)

	uViper.OnConfigChange(func(in fsnotify.Event) {
		reloadConfig(false)
		uCtxCancel()
	})
	uViper.WatchConfig()
}

// reloadConfig parses the config file and makes a new [tgmessenger.Messenger] instance, only validating at startup.
// reloadConfig panics on error.
func reloadConfig(validateToken bool) {
	err := uViper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err = uViper.Unmarshal(&config); err != nil {
		panic(err)
	}
	if messenger, err = tgmessenger.NewMessenger(config.BotToken, config.ChatId, config.TopicId, tgmessenger.ParseHTML, validateToken); err != nil {
		panic(err)
	}
}
