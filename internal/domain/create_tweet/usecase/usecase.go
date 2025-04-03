package usecase

import (
	"context"
	"uala-challenge/internal/domain/create_tweet"
	"uala-challenge/internal/model"
	"uala-challenge/utils/logger"
)

//go:generate mockery --name=tweetRepository --structname=TweetRepository --output=./mocks
type tweetRepository interface {
	CreateTweet(ctx context.Context, tweet *model.Tweet) error
}

//go:generate mockery --name=eventPublisher --structname=EventPublisher --output=./mocks
type eventPublisher interface {
	Publish(event *model.TweetCreatedEvent) error
}

type createTweetUseCase struct {
	tweetRepository tweetRepository
	eventPublisher  eventPublisher
}

func NewCreateTweetUseCase(tweetRepository tweetRepository, publisher eventPublisher) *createTweetUseCase {
	return &createTweetUseCase{
		tweetRepository: tweetRepository,
		eventPublisher:  publisher,
	}
}

func (useCase *createTweetUseCase) CreateTweet(ctx context.Context, tweet *model.Tweet) error {
	err := useCase.tweetRepository.CreateTweet(ctx, tweet)
	if err != nil {
		logger.GetLogger().Errorf("error creating tweet, message: %s, user_id: %s, err: %s", tweet.Content, tweet.UserID, err.Error())
		return createtweet.ErrInternal
	}

	pubErr := useCase.eventPublisher.Publish(model.NewTweetCreatedEvent(tweet))
	if pubErr != nil {
		logger.GetLogger().Errorf("error creating tweet creation event, message: %s, user_id: %s, err: %s", tweet.Content, tweet.UserID, err.Error())
	}

	return nil
}
