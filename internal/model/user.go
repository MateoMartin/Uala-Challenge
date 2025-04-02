package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        string   `json:"id"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}

func (u *User) AddFollower(followerId string) error {
	exist := false
	for _, v := range u.Followers {
		if v == followerId {
			exist = true
			break
		}
	}
	if !exist {
		u.Followers = append(u.Followers, followerId)
	}
	return nil
}

func (u *User) AddFollowing(followingId string) error {
	exist := false
	for _, v := range u.Following {
		if v == followingId {
			exist = true
			break
		}
	}
	if !exist {
		u.Following = append(u.Following, followingId)
	}

	return nil
}

func NewUser() *User {
	return &User{
		ID:        uuid.NewString(),
		Followers: make([]string, 0),
		Following: make([]string, 0),
	}
}
