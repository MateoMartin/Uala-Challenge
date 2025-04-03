package model

import "time"

type TweetCreatedEvent struct {
	ID          string     `json:"id"`
	Content     string     `json:"content"`
	UserID      string     `json:"user_id"`
	DateCreated *time.Time `json:"date_created"`
}

func (event *TweetCreatedEvent) ToTweet() *Tweet {
	return &Tweet{
		ID:          event.ID,
		Content:     event.Content,
		UserID:      event.UserID,
		DateCreated: event.DateCreated,
	}
}

func NewTweetCreatedEvent(tweet *Tweet) *TweetCreatedEvent {
	return &TweetCreatedEvent{
		ID:          tweet.ID,
		UserID:      tweet.UserID,
		Content:     tweet.Content,
		DateCreated: tweet.DateCreated,
	}
}
