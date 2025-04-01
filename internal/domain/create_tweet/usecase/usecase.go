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
type createTweetUseCase struct {
	tweetRepository tweetRepository
}

func NewCreateTweetUseCase(tweetRepository tweetRepository) *createTweetUseCase {
	return &createTweetUseCase{
		tweetRepository: tweetRepository,
	}
}

func (useCase *createTweetUseCase) CreateTweet(ctx context.Context, tweet *model.Tweet) error {
	err := useCase.tweetRepository.CreateTweet(ctx, tweet)
	if err != nil {
		logger.GetLogger().Errorf("error creating tweet, message: %s, user_id: %s, err: %s", tweet.Content, tweet.UserID, err.Error())
		return createtweet.ErrInternal
	}
	return nil
}
