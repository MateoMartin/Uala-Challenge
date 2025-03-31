package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

const (
	defaultPort = "8080"
)

type Config struct {
	Port               string `env:"PORT,default=8080"`
	GetStatusTimeout   int    `env:"GET_STATUS_TIMEOUT,default=1000"`
	CreateTweetTimeout int    `env:"CREATE_TWEET_TIMEOUT,default=1000"`
	FollowUserTimeout  int    `env:"FOLLOW_USER_TIMEOUT,default=1000"`
	GetTimelineTimeout int    `env:"GET_TIMELINE_TIMEOUT,default=1000"`
}

func LoadConfig() *Config {
	var config Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		log.Fatalf("Failed to process config: %s", err)
	}

	return &config
}
