package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"uala-challenge/internal/model"
)

func TestNewTweetDTOFromModelAndToModel(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name  string
		tweet *model.Tweet
	}{
		{
			name: "convert tweet to dto",
			tweet: &model.Tweet{
				ID:          "1",
				UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
				Content:     "test",
				DateCreated: &now,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dto := newTweetDTOFromModel(tc.tweet)
			assert.Equal(t, tc.tweet.ID, dto.ID)
			assert.Equal(t, tc.tweet.UserID, dto.UserID)
			assert.Equal(t, tc.tweet.Content, dto.Content)
			assert.Equal(t, tc.tweet.DateCreated, dto.DateCreated)

			convertedTweet := dto.toModel()
			assert.Equal(t, tc.tweet, convertedTweet)
		})
	}
}

func TestNewTimelineDTOFromModelAndToModel(t *testing.T) {
	now := time.Now()

	tweet1 := &model.Tweet{
		ID:          "1",
		UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
		Content:     "test",
		DateCreated: &now,
	}
	tweet2 := &model.Tweet{
		ID:          "2",
		UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
		Content:     "test",
		DateCreated: &now,
	}

	tests := []struct {
		name     string
		timeline *model.Timeline
	}{
		{
			name: "Timeline with two tweets",
			timeline: &model.Timeline{
				UserID: "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
				Tweets: []*model.Tweet{tweet1, tweet2},
			},
		},
		{
			name: "Timeline with no tweets",
			timeline: &model.Timeline{
				UserID: "fecf3e1f-f9d1-4976-9788-d3609504dda2",
				Tweets: []*model.Tweet{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dto := newTimelineDTOFromModel(tc.timeline)
			assert.Equal(t, tc.timeline.UserID, dto.UserID)
			assert.Equal(t, len(tc.timeline.Tweets), len(dto.Tweets))

			for i, tweetDTO := range dto.Tweets {
				expectedTweet := tc.timeline.Tweets[i]
				assert.Equal(t, expectedTweet.ID, tweetDTO.ID)
				assert.Equal(t, expectedTweet.UserID, tweetDTO.UserID)
				assert.Equal(t, expectedTweet.Content, tweetDTO.Content)
				assert.Equal(t, expectedTweet.DateCreated, tweetDTO.DateCreated)
			}

			convertedTimeline := dto.toModel()
			assert.Equal(t, tc.timeline, convertedTimeline)
		})
	}
}
