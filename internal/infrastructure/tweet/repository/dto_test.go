package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"uala-challenge/internal/model"
)

func TestNewDTOFromModel(t *testing.T) {
	now := time.Now()
	tweet := &model.Tweet{
		ID:          "dbc8370b-1e7d-4d1b-a131-fb0046ec9929",
		UserID:      "73790ac3-5d23-4131-b2f9-4ad1e3da61f6",
		Content:     "test",
		DateCreated: &now,
	}

	dto := newDTOFromModel(tweet)

	assert.Equal(t, tweet.ID, dto.ID)
	assert.Equal(t, tweet.UserID, dto.UserID)
	assert.Equal(t, tweet.Content, dto.Content)
	assert.Equal(t, tweet.DateCreated, dto.DateCreated)
}

func TestTweetDTO_ToModel(t *testing.T) {
	now := time.Now()
	dto := &tweetDTO{
		ID:          "1",
		UserID:      "user123",
		Content:     "Mensaje de prueba",
		DateCreated: &now,
	}

	tweet := dto.toModel()

	assert.Equal(t, dto.ID, tweet.ID)
	assert.Equal(t, dto.UserID, tweet.UserID)
	assert.Equal(t, dto.Content, tweet.Content)
	assert.Equal(t, dto.DateCreated, tweet.DateCreated)
}
