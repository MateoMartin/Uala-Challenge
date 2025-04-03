package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeline(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		expectedUserID string
	}{
		{
			name:           "Valid user timeline",
			userID:         "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
			expectedUserID: "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
		},
		{
			name:           "Empty userID timeline",
			userID:         "",
			expectedUserID: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			timeline := NewTimeline(tc.userID)
			assert.NotNil(t, timeline)
			assert.Equal(t, tc.expectedUserID, timeline.UserID)
			assert.Empty(t, timeline.Tweets)
		})
	}
}

func TestTimeline_AddTweet(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name          string
		tweetsToAdd   []*Tweet
		expectedCount int
	}{
		{
			name: "Add single tweet",
			tweetsToAdd: []*Tweet{
				{ID: "1", Content: "Hello world!", UserID: "0ebc1f72-b924-4d6e-bf35-885f50f5034a", DateCreated: &now},
			},
			expectedCount: 1,
		},
		{
			name: "Add multiple tweets",
			tweetsToAdd: []*Tweet{
				{ID: "1", Content: "Hello world!", UserID: "0ebc1f72-b924-4d6e-bf35-885f50f5034a", DateCreated: &now},
				{ID: "2", Content: "Another tweet", UserID: "0ebc1f72-b924-4d6e-bf35-885f50f5034a", DateCreated: &now},
			},
			expectedCount: 2,
		},
		{
			name:          "Add no tweets",
			tweetsToAdd:   []*Tweet{},
			expectedCount: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			timeline := NewTimeline("0ebc1f72-b924-4d6e-bf35-885f50f5034a")
			for _, tweet := range tc.tweetsToAdd {
				timeline.AddTweet(tweet)
			}
			assert.Equal(t, tc.expectedCount, len(timeline.Tweets))
			for i, tweet := range tc.tweetsToAdd {
				assert.Equal(t, tweet, timeline.Tweets[i])
			}
		})
	}
}
