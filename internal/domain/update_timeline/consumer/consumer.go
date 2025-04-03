package consumer

import (
	"context"
	"time"
	"uala-challenge/internal/model"
	"uala-challenge/utils/logger"
)

//go:generate mockery --name=followRepository --structname=FollowRepository --output=./mocks
type followRepository interface {
	GetUser(ctx context.Context, userID string) (*model.User, error)
}

//go:generate mockery --name=timelineRepository --structname=TimelineRepository --output=./mocks
type timelineRepository interface {
	UpdateTimelinesForTweet(ctx context.Context, tweet *model.Tweet, followers []string) error
}

type updateTimelineConsumer struct {
	updateMaxDuration  time.Duration
	eventChan          <-chan *model.TweetCreatedEvent
	timelineRepository timelineRepository
	followRepository   followRepository
}

func NewUpdateTimelineConsumer(eventChan <-chan *model.TweetCreatedEvent, timelineRepository timelineRepository, followRepository followRepository, maxUpdateDuration time.Duration) *updateTimelineConsumer {
	return &updateTimelineConsumer{
		eventChan:          eventChan,
		timelineRepository: timelineRepository,
		followRepository:   followRepository,
		updateMaxDuration:  maxUpdateDuration,
	}
}

func (c *updateTimelineConsumer) Start() {
	go func() {
		for {
			select {
			case event := <-c.eventChan:
				c.processEvent(event)
			}
		}
	}()
}

func (c *updateTimelineConsumer) processEvent(event *model.TweetCreatedEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*c.updateMaxDuration)
	defer cancel()

	user, err := c.followRepository.GetUser(ctx, event.UserID)
	if err != nil {
		logger.GetLogger().Errorf("error getting user to update timeline, user_id: %s, error: %s", event.UserID, err.Error())
		return
	}

	err = c.timelineRepository.UpdateTimelinesForTweet(ctx, event.ToTweet(), user.Followers)
	if err != nil {
		logger.GetLogger().Errorf("error updating timeline for user, user_id: %s, error: %s", event.UserID, err.Error())
	}

	return
}
