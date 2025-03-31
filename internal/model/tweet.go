package model

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	MaxCharacters = 280
)

var (
	ErrorMaxCharacters = fmt.Errorf("tweet should have less than %d characters", MaxCharacters)
)

type Tweet struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

func NewTweet(userID, Content string) (*Tweet, error) {
	if len(Content) > MaxCharacters {
		return nil, ErrorMaxCharacters
	}
	return &Tweet{
		ID:      uuid.NewString(),
		UserID:  userID,
		Content: Content,
	}, nil
}
