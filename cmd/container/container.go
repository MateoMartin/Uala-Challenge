package container

import (
	"github.com/gin-gonic/gin"
	createTweetHandler "uala-challenge/internal/domain/create_tweet/handler"
	createTweetUseCase "uala-challenge/internal/domain/create_tweet/usecase"
	followUserHandler "uala-challenge/internal/domain/follow_user/handler"
	followUserUseCase "uala-challenge/internal/domain/follow_user/usecase"
	getStatusHandler "uala-challenge/internal/domain/get_status/handler"
	tweetRepository "uala-challenge/internal/infrastructure/tweet/repository"
	userRepository "uala-challenge/internal/infrastructure/user/repository"
)

type container struct {
	GetStatusHandler   gin.HandlerFunc
	CreateTweetHandler gin.HandlerFunc
	FollowUserHandler  gin.HandlerFunc
}

func LoadContainer() *container {
	tweetRepository := tweetRepository.NewInMemoryTweetRepository()
	userRepository := userRepository.NewInMemoryUserRepository()

	createTweetUseCase := createTweetUseCase.NewCreateTweetUseCase(tweetRepository)
	followUserUseCase := followUserUseCase.NewFollowUserUseCase(userRepository)

	getStatusHandler := getStatusHandler.NewGetStatusHandler()
	createTweetHandler := createTweetHandler.NewCreateTweetHandler(createTweetUseCase)
	followUserTweetHandler := followUserHandler.NewFollowUserHandler(followUserUseCase)

	return &container{
		GetStatusHandler:   getStatusHandler.Handle,
		CreateTweetHandler: createTweetHandler.Handle,
		FollowUserHandler:  followUserTweetHandler.Handle,
	}
}
