package repository

import (
	"context"
	"uala-challenge/internal/model"
)

type inMemoryTweetRepository struct {
	tweetByID map[string]*tweetDTO
}

func NewInMemoryTweetRepository() *inMemoryTweetRepository {
	tweetByID := make(map[string]*tweetDTO, 0)

	return &inMemoryTweetRepository{
		tweetByID: tweetByID,
	}
}

func (r *inMemoryTweetRepository) CreateTweet(ctx context.Context, tweet *model.Tweet) error {
	r.tweetByID[tweet.ID] = newDTOFromModel(tweet)
	return nil
}
