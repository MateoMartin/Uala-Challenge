package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name              string
		expectedFollowing []string
		expectedFollowers []string
	}{
		{
			name:              "user valid",
			expectedFollowers: []string{},
			expectedFollowing: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := NewUser()
			assert.NotNil(t, user)
			assert.Equal(t, tt.expectedFollowers, user.Followers)
			assert.NotNil(t, user.Following)
			assert.Equal(t, tt.expectedFollowing, user.Following)
		})
	}
}

func TestAddFollower(t *testing.T) {
	tests := []struct {
		name       string
		user       *User
		followerId string
	}{
		{
			name:       "Add follower successfully",
			user:       NewUser(),
			followerId: "5124e15c-7806-4c65-9207-e71a9c384754",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := NewUser()
			user.AddFollower(tt.followerId)
			assert.NotNil(t, user)
			assert.NotNil(t, user.Followers)
			assert.Equal(t, tt.followerId, user.Followers[0])
		})
	}
}

func TestAddFolling(t *testing.T) {
	tests := []struct {
		name        string
		user        *User
		followingId string
	}{
		{
			name:        "Add following successfully",
			user:        NewUser(),
			followingId: "5124e15c-7806-4c65-9207-e71a9c384754",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := NewUser()
			user.AddFollowing(tt.followingId)
			assert.NotNil(t, user)
			assert.NotNil(t, user.Following)
			assert.Equal(t, tt.followingId, user.Following[0])
		})
	}
}
