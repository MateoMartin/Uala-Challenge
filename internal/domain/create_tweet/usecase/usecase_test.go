package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"uala-challenge/internal/domain/create_tweet"
	"uala-challenge/internal/domain/create_tweet/usecase/mocks"
	"uala-challenge/internal/model"
)

func TestCreateTweetUseCase_CreateTweet(t *testing.T) {
	tests := []struct {
		name          string
		tweet         *model.Tweet
		setupMocks    func(repo *mocks.TweetRepository)
		expectedError error
	}{
		{
			name:  "should return error if repository fails",
			tweet: &model.Tweet{ID: "1", UserID: "dbc8370b-1e7d-4d1b-a131-fb0046ec9929", Content: "test"},
			setupMocks: func(repo *mocks.TweetRepository) {
				repo.On("CreateTweet", mock.Anything, mock.Anything).
					Return(errors.New("repository error")).Once()
			},
			expectedError: createtweet.ErrInternal,
		},
		{
			name:  "should create tweet successfully",
			tweet: &model.Tweet{ID: "1", UserID: "dbc8370b-1e7d-4d1b-a131-fb0046ec9929", Content: "test"},
			setupMocks: func(repo *mocks.TweetRepository) {
				repo.On("CreateTweet", mock.Anything, mock.Anything).
					Return(nil).Once()
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mocks.TweetRepository)
			publisher := new(mocks.EventPublisher)
			publisher.On("Publish", mock.Anything).Return(nil)
			tt.setupMocks(repo)

			useCase := NewCreateTweetUseCase(repo, publisher)
			err := useCase.CreateTweet(context.Background(), tt.tweet)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
			repo.AssertExpectations(t)
		})
	}
}
