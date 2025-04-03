package container

import (
	"github.com/gin-gonic/gin"
	createTweetHandler "uala-challenge/internal/domain/create_tweet/handler"
	createTweetUseCase "uala-challenge/internal/domain/create_tweet/usecase"
	followUserHandler "uala-challenge/internal/domain/follow_user/handler"
	followUserUseCase "uala-challenge/internal/domain/follow_user/usecase"
	getStatusHandler "uala-challenge/internal/domain/get_status/handler"
	timelineConsumer "uala-challenge/internal/domain/update_timeline/consumer"
	eventPublisher "uala-challenge/internal/infrastructure/event"
	timelineRepository "uala-challenge/internal/infrastructure/timeline/repository"
	tweetRepository "uala-challenge/internal/infrastructure/tweet/repository"
	userRepository "uala-challenge/internal/infrastructure/user/repository"
	"uala-challenge/internal/model"
)

type container struct {
	GetStatusHandler   gin.HandlerFunc
	CreateTweetHandler gin.HandlerFunc
	FollowUserHandler  gin.HandlerFunc
}

func LoadContainer() *container {
	tweetRepository := tweetRepository.NewInMemoryTweetRepository()
	userRepository := userRepository.NewInMemoryUserRepository()
	timelineRepository := timelineRepository.NewInMemoryTimelineRepository()

	eventChannel := make(chan *model.TweetCreatedEvent)
	eventPublisher := eventPublisher.NewInMemoryEventPublisher(eventChannel)

	createTweetUseCase := createTweetUseCase.NewCreateTweetUseCase(tweetRepository, eventPublisher)
	followUserUseCase := followUserUseCase.NewFollowUserUseCase(userRepository)

	getStatusHandler := getStatusHandler.NewGetStatusHandler()
	createTweetHandler := createTweetHandler.NewCreateTweetHandler(createTweetUseCase)
	followUserTweetHandler := followUserHandler.NewFollowUserHandler(followUserUseCase)

	timelineUpdaterConsumer := timelineConsumer.NewUpdateTimelineConsumer(eventChannel, timelineRepository, userRepository, 5000)
	timelineUpdaterConsumer.Start()

	return &container{
		GetStatusHandler:   getStatusHandler.Handle,
		CreateTweetHandler: createTweetHandler.Handle,
		FollowUserHandler:  followUserTweetHandler.Handle,
	}
}
