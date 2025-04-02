package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"uala-challenge/internal/domain/follow_user"
	"uala-challenge/internal/domain/follow_user/usecase/mocks"
)

func TestFollowUserUseCase_FollowUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		userIDToFollow string
		setupMocks     func(repo *mocks.UserRepository)
		expectedError  error
	}{
		{
			name:           "should return error if repository fails",
			userID:         "dbc8370b-1e7d-4d1b-a131-fb0046ec9929",
			userIDToFollow: "ec04f84a-8467-4db3-b4d9-5a97c7e59349",
			setupMocks: func(repo *mocks.UserRepository) {
				repo.On("FollowUser", mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("repository error")).Once()
			},
			expectedError: followuser.ErrInternal,
		},
		{
			name:           "should follow user successfully",
			userID:         "dbc8370b-1e7d-4d1b-a131-fb0046ec9929",
			userIDToFollow: "ec04f84a-8467-4db3-b4d9-5a97c7e59349",
			setupMocks: func(repo *mocks.UserRepository) {
				repo.On("FollowUser", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).Once()
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.UserRepository)
			tt.setupMocks(repo)

			useCase := NewFollowUserUseCase(repo)
			err := useCase.FollowUser(context.Background(), tt.userID, tt.userIDToFollow)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			repo.AssertExpectations(t)
		})
	}
}
