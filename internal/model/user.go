package model

import "github.com/google/uuid"

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Following []string `json:"following"`
}

func NewUser(username string) *User {
	return &User{
		ID:        uuid.NewString(),
		Username:  username,
		Following: make([]string, 0),
	}
}
