package repository

import (
	"time"
	"uala-challenge/internal/model"
)

type tweetDTO struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Content     string     `json:"content"`
	DateCreated *time.Time `json:"date_created"`
}

func newDTOFromModel(tweet *model.Tweet) *tweetDTO {
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
