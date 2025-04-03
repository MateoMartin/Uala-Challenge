package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTweetCreatedEvent(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name          string
		inputTweet    *Tweet
		expectedEvent *TweetCreatedEvent
	}{
		{
			name: "event created sucessfully",
			inputTweet: &Tweet{
				ID:          "1",
				Content:     "test",
				UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
				DateCreated: &now,
			},
			expectedEvent: &TweetCreatedEvent{
				ID:          "1",
				Content:     "test",
				UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
				DateCreated: &now,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			event := NewTweetCreatedEvent(tc.inputTweet)
			assert.Equal(t, tc.expectedEvent.ID, event.ID)
			assert.Equal(t, tc.expectedEvent.Content, event.Content)
			assert.Equal(t, tc.expectedEvent.UserID, event.UserID)
			assert.Equal(t, tc.expectedEvent.DateCreated, event.DateCreated)
		})
	}
}

func TestTweetCreatedEvent_ToTweet(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name          string
		inputEvent    *TweetCreatedEvent
		expectedTweet *Tweet
	}{
		{
			name: "conversion to tweet sucessfully",
			inputEvent: &TweetCreatedEvent{
				ID:          "1",
				Content:     "test",
				UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
				DateCreated: &now,
			},
			expectedTweet: &Tweet{
				ID:          "1",
				Content:     "test",
				UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
				DateCreated: &now,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tweet := tc.inputEvent.ToTweet()
			assert.Equal(t, tc.expectedTweet.ID, tweet.ID)
			assert.Equal(t, tc.expectedTweet.Content, tweet.Content)
			assert.Equal(t, tc.expectedTweet.UserID, tweet.UserID)
			assert.Equal(t, tc.expectedTweet.DateCreated, tweet.DateCreated)
		})
	}
}
