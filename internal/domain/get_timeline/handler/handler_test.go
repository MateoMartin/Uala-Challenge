package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"uala-challenge/internal/domain/get_timeline/handler/mocks"
	"uala-challenge/internal/model"
)

func TestGetTimelineHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name                string
		userID              string
		mockGetTimelineFunc func(ctx context.Context, userID string) (*model.Timeline, error)
		expectedStatus      int
		expectedBodySubstr  string
	}{
		{
			name:   "UseCase returns error",
			userID: "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
			mockGetTimelineFunc: func(ctx context.Context, userID string) (*model.Timeline, error) {
				return nil, errors.New("internal error")
			},
			expectedStatus:     http.StatusInternalServerError,
			expectedBodySubstr: "internal error",
		},
		{
			name:   "Successful timeline retrieval",
			userID: "fd49bd9c-c552-426a-8fdb-32c51ab6f8f6",
			mockGetTimelineFunc: func(ctx context.Context, userID string) (*model.Timeline, error) {
				return &model.Timeline{
					UserID: userID,
					Tweets: []*model.Tweet{},
				}, nil
			},
			expectedStatus:     http.StatusOK,
			expectedBodySubstr: `"user_id":"fd49bd9c-c552-426a-8fdb-32c51ab6f8f6"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUseCase := new(mocks.GetTimelineUseCase)
			if tc.mockGetTimelineFunc != nil {
				mockUseCase.
					On("GetTimeline", mock.Anything, tc.userID).
					Return(tc.mockGetTimelineFunc(context.Background(), tc.userID))
			}

			h := NewGetTimelineHandler(mockUseCase)

			router := gin.Default()
			router.GET("/timeline/:user_id", h.Handle)

			reqPath := "/timeline/" + tc.userID
			req, err := http.NewRequest(http.MethodGet, reqPath, nil)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tc.expectedStatus, recorder.Code)
			
			mockUseCase.AssertExpectations(t)
		})
	}
}
