package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"uala-challenge/internal/model"
)

func TestNewInMemoryTweetRepository(t *testing.T) {
	repo := NewInMemoryTweetRepository()
	assert.NotNil(t, repo)
	assert.Empty(t, repo.tweetByID)
}

func TestInMemoryTweetRepository_CreateTweet(t *testing.T) {
	repo := NewInMemoryTweetRepository()
	ctx := context.TODO()

	tweet := &model.Tweet{
		ID: "123",
	}

	err := repo.CreateTweet(ctx, tweet)
	assert.NoError(t, err)

	storedTweet, ok := repo.tweetByID[tweet.ID]
	assert.True(t, ok)
	assert.Equal(t, tweet.ID, storedTweet.ID)
}
