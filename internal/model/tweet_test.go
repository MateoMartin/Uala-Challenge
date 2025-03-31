package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTweet(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		content string
		wantErr error
	}{
		{
			name:    "Valid tweet with short content should be valid",
			userID:  "61112ed1-d78a-4b47-8b54-fcdb0f68f476",
			content: "Hello, world!",
			wantErr: nil,
		},
		{
			name:    "Valid tweet with content exactly at max length should be valid",
			userID:  "61112ed1-d78a-4b47-8b54-fcdb0f68f476",
			content: strings.Repeat("a", MaxCharacters),
			wantErr: nil,
		},
		{
			name:    "Invalid tweet exceeding max characters should return error",
			userID:  "61112ed1-d78a-4b47-8b54-fcdb0f68f476",
			content: strings.Repeat("a", MaxCharacters+1),
			wantErr: ErrorMaxCharacters,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tweet, err := NewTweet(tt.userID, tt.content)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, tweet)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, tweet)
				assert.Equal(t, tt.userID, tweet.UserID)
				assert.Equal(t, tt.content, tweet.Content)
				assert.NotEmpty(t, tweet.ID)
			}
		})
	}
}
