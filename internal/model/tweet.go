package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	MaxCharacters = 280
)

var (
	ErrorMaxCharacters = fmt.Errorf("tweet should have less than %d characters", MaxCharacters)
)

type Tweet struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Content     string     `json:"content"`
	DateCreated *time.Time `json:"date_created"`
}

func NewTweet(userID, Content string) (*Tweet, error) {
	if len(Content) > MaxCharacters {
		return nil, ErrorMaxCharacters
	}

	dateCreated := time.Now()
	return &Tweet{
		ID:          uuid.NewString(),
		UserID:      userID,
		Content:     Content,
		DateCreated: &dateCreated,
	}, nil
}
