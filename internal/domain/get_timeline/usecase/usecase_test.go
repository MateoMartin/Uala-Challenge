package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"uala-challenge/internal/domain/get_timeline/usecase/mocks"
	"uala-challenge/internal/model"
)

func TestGetTimeline(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name             string
		userID           string
		expectedTimeline *model.Timeline
		expectedError    error
		setupMock        func(m *mocks.TimelineRepository)
	}{
		{
			name:   "Successful timeline retrieval",
			userID: "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
			expectedTimeline: &model.Timeline{
				UserID: "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
				Tweets: []*model.Tweet{
					{
						ID:          "tweet1",
						Content:     "content",
						UserID:      "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
						DateCreated: &now,
					},
				},
			},
			expectedError: nil,
			setupMock: func(m *mocks.TimelineRepository) {
				m.
					On("GetTimelineByUserID", mock.Anything, "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6").
					Return(&model.Timeline{
						UserID: "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
						Tweets: []*model.Tweet{
							{
								ID:          "tweet1",
								Content:     "test",
								UserID:      "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
								DateCreated: &now,
							},
						},
					}, nil)
			},
		},
		{
			name:             "Timeline repository error",
			userID:           "948143c7-590b-4667-abf6-94389910700c",
			expectedTimeline: nil,
			expectedError:    errors.New("repo error"),
			setupMock: func(m *mocks.TimelineRepository) {
				m.
					On("GetTimelineByUserID", mock.Anything, "948143c7-590b-4667-abf6-94389910700c").
					Return(nil, errors.New("repo error"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.TimelineRepository)
			tc.setupMock(mockRepo)

			useCaseInstance := NewTimelineUseCase(mockRepo)
			timeline, err := useCaseInstance.GetTimeline(context.Background(), tc.userID)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
				assert.Nil(t, timeline)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tc.expectedTimeline.Tweets), len(timeline.Tweets))
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
