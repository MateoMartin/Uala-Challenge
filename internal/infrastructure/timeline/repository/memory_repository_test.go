package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"uala-challenge/internal/model"
)

func TestUpdateTimelinesForTweet(t *testing.T) {
	now := time.Now()
	testTweet := &model.Tweet{
		ID:          "tweet1",
		Content:     "test",
		UserID:      "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
		DateCreated: &now,
	}

	tests := []struct {
		name          string
		followers     []string
		preSetup      func(repo *inMemoryTimelineRepository)
		expectedCount map[string]int
	}{
		{
			name:          "No followers provided",
			followers:     []string{},
			preSetup:      nil,
			expectedCount: map[string]int{},
		},
		{
			name:      "Single follower, new timeline",
			followers: []string{"follower1"},
			preSetup:  nil,
			expectedCount: map[string]int{
				"follower1": 1,
			},
		},
		{
			name:      "Multiple followers, new timelines",
			followers: []string{"follower1", "follower2"},
			preSetup:  nil,
			expectedCount: map[string]int{
				"follower1": 1,
				"follower2": 1,
			},
		},
		{
			name:      "Existing timeline gets appended",
			followers: []string{"follower1"},
			preSetup: func(repo *inMemoryTimelineRepository) {
				initialTweet := &model.Tweet{
					ID:          "1",
					Content:     "test",
					UserID:      "follower1",
					DateCreated: &now,
				}
				timeline := model.NewTimeline("follower1")
				timeline.AddTweet(initialTweet)
				repo.timelineByUserID["follower1"] = newTimelineDTOFromModel(timeline)
			},
			expectedCount: map[string]int{
				"follower1": 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewInMemoryTimelineRepository()
			if tc.preSetup != nil {
				tc.preSetup(repo)
			}

			err := repo.UpdateTimelinesForTweet(context.Background(), testTweet, tc.followers)
			assert.NoError(t, err)

			for follower, expectedNum := range tc.expectedCount {
				dto, exists := repo.timelineByUserID[follower]
				assert.True(t, exists)
				timeline := dto.toModel()
				assert.Equal(t, expectedNum, len(timeline.Tweets))
				if expectedNum > 0 {
					lastTweet := timeline.Tweets[len(timeline.Tweets)-1]
					assert.Equal(t, testTweet.ID, lastTweet.ID)
				}
			}

			if len(tc.followers) == 0 {
				assert.Empty(t, repo.timelineByUserID)
			}
		})
	}
}

func TestGetTimelineByUserID(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name           string
		userID         string
		preSetup       func(repo *inMemoryTimelineRepository)
		expectedTweets int
	}{
		{
			name:   "Timeline exists with tweets",
			userID: "follower1",
			preSetup: func(repo *inMemoryTimelineRepository) {
				timeline := model.NewTimeline("follower1")
				tweet := &model.Tweet{
					ID:          "tweet1",
					Content:     "test",
					UserID:      "follower1",
					DateCreated: &now,
				}
				timeline.AddTweet(tweet)
				repo.timelineByUserID["follower1"] = newTimelineDTOFromModel(timeline)
			},
			expectedTweets: 1,
		},
		{
			name:           "Timeline does not exist returns new empty timeline",
			userID:         "newUser",
			preSetup:       func(repo *inMemoryTimelineRepository) {},
			expectedTweets: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewInMemoryTimelineRepository()
			if tc.preSetup != nil {
				tc.preSetup(repo)
			}

			timeline, err := repo.GetTimelineByUserID(context.Background(), tc.userID)
			assert.NoError(t, err)
			assert.Equal(t, tc.userID, timeline.UserID)
			assert.Equal(t, tc.expectedTweets, len(timeline.Tweets), tc.expectedTweets)
		})
	}
}
