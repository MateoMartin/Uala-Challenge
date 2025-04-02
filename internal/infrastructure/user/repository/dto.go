package repository

import (
	"uala-challenge/internal/model"
)

type userDTO struct {
	ID        string   `json:"id"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}

func newDTOFromModel(user *model.User) *userDTO {
	return &userDTO{
		ID:        user.ID,
		Following: user.Following,
		Followers: user.Followers,
	}
}

func (u *userDTO) toModel() *model.User {
	return &model.User{
		ID:        u.ID,
		Followers: u.Followers,
		Following: u.Following,
	}
}
