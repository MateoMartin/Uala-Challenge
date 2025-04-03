package main

import (
	"uala-challenge/cmd/container"
	"uala-challenge/config"
	"uala-challenge/router"
	"uala-challenge/utils/logger"
)

// @title Uala Challenge API
// @version 1.0
// @description Twitter service for Uala challenge
// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.LoadConfig()
	container := container.LoadContainer()
	router := router.SetupRouter(
		cfg,
		container.GetStatusHandler,
		container.CreateTweetHandler,
		container.FollowUserHandler,
		container.GetTimelineHandler,
	)

	logger.GetLogger().Infof("Service is starting at port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
