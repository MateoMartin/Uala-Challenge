package consumer

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"uala-challenge/internal/domain/update_timeline/consumer/mocks"
	"uala-challenge/internal/model"
)

func TestUpdateTimelineConsumer_ProcessEvent(t *testing.T) {
	now := time.Now()
	testTweet := &model.Tweet{
		ID:          "1",
		Content:     "test",
		UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
		DateCreated: &now,
	}
	event := &model.TweetCreatedEvent{
		ID:          "1",
		Content:     "test",
		UserID:      "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
		DateCreated: &now,
	}

	tests := []struct {
		name                 string
		getUserError         error
		updateTimelineError  error
		expectedUpdateCalled bool
	}{
		{
			name:                 "process event successfully",
			getUserError:         nil,
			updateTimelineError:  nil,
			expectedUpdateCalled: true,
		},
		{
			name:                 "error getting user",
			getUserError:         errors.New("error getting user"),
			updateTimelineError:  nil,
			expectedUpdateCalled: false,
		},
		{
			name:                 "error updating timeline",
			getUserError:         nil,
			updateTimelineError:  errors.New("failed update"),
			expectedUpdateCalled: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			followRepoMock := new(mocks.FollowRepository)
			if tc.getUserError != nil {
				followRepoMock.
					On("GetUser", mock.Anything, event.UserID).
					Return((*model.User)(nil), tc.getUserError)
			} else {
				user := &model.User{
					ID:        "0ebc1f72-b924-4d6e-bf35-885f50f5034a",
					Followers: []string{"1", "2"},
				}
				followRepoMock.
					On("GetUser", mock.Anything, event.UserID).
					Return(user, nil)
			}

			timelineRepoMock := new(mocks.TimelineRepository)
			if tc.expectedUpdateCalled {
				timelineRepoMock.
					On("UpdateTimelinesForTweet", mock.Anything, testTweet, mock.Anything).
					Return(tc.updateTimelineError)
			}

			consumerInstance := NewUpdateTimelineConsumer(nil, timelineRepoMock, followRepoMock, 1000)

			consumerInstance.processEvent(event)

			followRepoMock.AssertExpectations(t)
			if tc.expectedUpdateCalled {
				timelineRepoMock.AssertExpectations(t)
			} else {
				timelineRepoMock.AssertNotCalled(t, "UpdateTimelinesForTweet", mock.Anything, mock.Anything, mock.Anything)
			}
		})
	}
}
