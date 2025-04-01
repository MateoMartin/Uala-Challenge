package container

import (
	"github.com/gin-gonic/gin"
	createTweetHandler "uala-challenge/internal/domain/create_tweet/handler"
	createTweetUseCase "uala-challenge/internal/domain/create_tweet/usecase"
	getStatusHandler "uala-challenge/internal/domain/get_status/handler"
	tweetRepository "uala-challenge/internal/infrastructure/tweet/repository"
)

type container struct {
	GetStatusHandler   gin.HandlerFunc
	CreateTweetHandler gin.HandlerFunc
}

func LoadContainer() *container {
	tweetRepository := tweetRepository.NewInMemoryTweetRepository()

	createTweetUseCase := createTweetUseCase.NewCreateTweetUseCase(tweetRepository)

	getStatusHandler := getStatusHandler.NewGetStatusHandler()
	createTweetHandler := createTweetHandler.NewCreateTweetHandler(createTweetUseCase)

	return &container{
		GetStatusHandler:   getStatusHandler.Handle,
		CreateTweetHandler: createTweetHandler.Handle,
	}
}
