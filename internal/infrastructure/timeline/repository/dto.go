package repository

import (
	"time"
	"uala-challenge/internal/model"
)

type timelineDTO struct {
	UserID string      `json:"user_id"`
	Tweets []*tweetDTO `json:"tweets"`
}

type tweetDTO struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Content     string     `json:"content"`
	DateCreated *time.Time `json:"date_created"`
}

func newTweetDTOFromModel(tweet *model.Tweet) *tweetDTO {
	return &tweetDTO{
		ID:          tweet.ID,
		UserID:      tweet.UserID,
		Content:     tweet.Content,
		DateCreated: tweet.DateCreated,
	}
}

func (t *tweetDTO) toModel() *model.Tweet {
	return &model.Tweet{
		ID:          t.ID,
		UserID:      t.UserID,
		Content:     t.Content,
		DateCreated: t.DateCreated,
	}
}

func newTimelineDTOFromModel(timeline *model.Timeline) *timelineDTO {
	tweetsDto := make([]*tweetDTO, len(timeline.Tweets))
	for i, tweet := range timeline.Tweets {
		tweetsDto[i] = newTweetDTOFromModel(tweet)
	}
	return &timelineDTO{
		UserID: timeline.UserID,
		Tweets: tweetsDto,
	}
}

func (t *timelineDTO) toModel() *model.Timeline {
	tweets := make([]*model.Tweet, len(t.Tweets))
	for i, tweet := range t.Tweets {
		tweets[i] = tweet.toModel()
	}
	return &model.Timeline{
		UserID: t.UserID,
		Tweets: tweets,
	}
}
