package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name              string
		username          string
		expectedUsername  string
		expectedFollowing []string
	}{
		{
			name:              "user with valid username",
			username:          "test",
			expectedUsername:  "test",
			expectedFollowing: []string{},
		},
		{
			name:              "user with empty username",
			username:          "",
			expectedUsername:  "",
			expectedFollowing: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := NewUser(tt.username)
			assert.NotNil(t, user)
			assert.Equal(t, tt.expectedUsername, user.Username)
			assert.NotNil(t, user.Following)
			assert.Equal(t, tt.expectedFollowing, user.Following)
		})
	}
}
