package repository

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"uala-challenge/internal/model"
)

func newTestUser(id string) *model.User {
	return &model.User{
		ID:        id,
		Following: []string{},
		Followers: []string{},
	}
}

func newUserDTO(id string) *userDTO {
	return newDTOFromModel(newTestUser(id))
}

func TestInMemoryUserRepository_GetUser(t *testing.T) {
	repo := &inMemoryUserRepository{
		usersByID: make(map[string]*userDTO),
		mutex:     sync.RWMutex{},
	}

	tests := []struct {
		name      string
		setup     func()
		userID    string
		expectErr bool
	}{
		{
			name: "returns user successfully",
			setup: func() {
				repo.usersByID["0f589c7a-e25a-4805-95e6-47c013e23bd3"] = newUserDTO("0f589c7a-e25a-4805-95e6-47c013e23bd3")
			},
			userID:    "0f589c7a-e25a-4805-95e6-47c013e23bd3",
			expectErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup()
			user, err := repo.GetUser(context.Background(), tc.userID)
			assert.NoError(t, err)
			assert.Equal(t, tc.userID, user.ID)
			assert.Empty(t, user.Following)
			assert.Empty(t, user.Followers)
		})
	}
}

func TestInMemoryUserRepository_FollowUser(t *testing.T) {
	tests := []struct {
		name              string
		initialUsers      map[string]*userDTO
		userID            string
		userIDToFollow    string
		expectedFollowing []string
		expectedFollowers []string
	}{
		{
			name: "Follow user use user in DB if exists",
			initialUsers: map[string]*userDTO{
				"0f589c7a-e25a-4805-95e6-47c013e23bd3": newUserDTO("0f589c7a-e25a-4805-95e6-47c013e23bd3"),
				"dd6858df-4786-48c5-b8e5-a674b0f191a1": newUserDTO("dd6858df-4786-48c5-b8e5-a674b0f191a1"),
			},
			userID:            "0f589c7a-e25a-4805-95e6-47c013e23bd3",
			userIDToFollow:    "dd6858df-4786-48c5-b8e5-a674b0f191a1",
			expectedFollowing: []string{"dd6858df-4786-48c5-b8e5-a674b0f191a1"},
			expectedFollowers: []string{"0f589c7a-e25a-4805-95e6-47c013e23bd3"},
		},
		{
			name: "Follow user creates user if it doesn't exist on DB",
			initialUsers: map[string]*userDTO{
				"dd6858df-4786-48c5-b8e5-a674b0f191a1": newUserDTO("dd6858df-4786-48c5-b8e5-a674b0f191a1"),
			},
			userID:            "0f589c7a-e25a-4805-95e6-47c013e23bd3",
			userIDToFollow:    "dd6858df-4786-48c5-b8e5-a674b0f191a1",
			expectedFollowing: []string{"dd6858df-4786-48c5-b8e5-a674b0f191a1"},
			expectedFollowers: []string{"0f589c7a-e25a-4805-95e6-47c013e23bd3"},
		},
		{
			name: "Follow user creates user if user to follow doesn't exist on DB",
			initialUsers: map[string]*userDTO{
				"0f589c7a-e25a-4805-95e6-47c013e23bd3": newUserDTO("0f589c7a-e25a-4805-95e6-47c013e23bd3"),
			},
			userID:            "0f589c7a-e25a-4805-95e6-47c013e23bd3",
			userIDToFollow:    "dd6858df-4786-48c5-b8e5-a674b0f191a1",
			expectedFollowing: []string{"dd6858df-4786-48c5-b8e5-a674b0f191a1"},
			expectedFollowers: []string{"0f589c7a-e25a-4805-95e6-47c013e23bd3"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &inMemoryUserRepository{
				usersByID: make(map[string]*userDTO),
				mutex:     sync.RWMutex{},
			}

			for id, dto := range tc.initialUsers {
				repo.usersByID[id] = dto
			}

			err := repo.FollowUser(context.Background(), tc.userID, tc.userIDToFollow)
			assert.NoError(t, err)

			user, err := repo.GetUser(context.Background(), tc.userID)
			assert.NoError(t, err)
			followedUser, err := repo.GetUser(context.Background(), tc.userIDToFollow)
			assert.NoError(t, err)

			assert.Contains(t, user.Following, tc.userIDToFollow)
			assert.Contains(t, followedUser.Followers, tc.userID)
		})
	}
}
