package repository

import (
	"context"
	"sync"
	"uala-challenge/internal/model"
)

const (
	maxTweetsOnTimeline = 100
)

type inMemoryTimelineRepository struct {
	timelineByUserID map[string]*timelineDTO
	mutex            sync.RWMutex
}

func NewInMemoryTimelineRepository() *inMemoryTimelineRepository {
	timelineByUserID := make(map[string]*timelineDTO, 0)

	return &inMemoryTimelineRepository{
		timelineByUserID: timelineByUserID,
	}
}

func (r *inMemoryTimelineRepository) UpdateTimelinesForTweet(ctx context.Context, tweet *model.Tweet, followers []string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, follower := range followers {
		timelineDTO := r.timelineByUserID[follower]
		if timelineDTO == nil {
			timelineDTO = newTimelineDTOFromModel(model.NewTimeline(follower))
		}
		timeline := timelineDTO.toModel()
		timeline.AddTweet(tweet)

		r.timelineByUserID[follower] = newTimelineDTOFromModel(timeline)
	}

	return nil
}

func (r *inMemoryTimelineRepository) GetTimelineByUserID(ctx context.Context, userID string) (*model.Timeline, error) {
	timelineDTO := r.timelineByUserID[userID]

	if timelineDTO == nil {
		return model.NewTimeline(userID), nil
	}
	return timelineDTO.toModel(), nil
}
