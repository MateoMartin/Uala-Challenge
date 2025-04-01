package handler

import "uala-challenge/internal/model"

type createTweetDTO struct {
	UserID  string `json:"user_id" binding:"required" example:"0f089136-3f38-4757-840c-d0c954782457"`
	Content string `json:"content" binding:"required" example:"Hello!"`
}

func (dto createTweetDTO) toTweet() (*model.Tweet, error) {
	return model.NewTweet(dto.UserID, dto.Content)
}
