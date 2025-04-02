package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"uala-challenge/internal/domain/follow_user/handler/mocks"
)

func TestFollowUserHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name               string
		requestBody        string
		mockUseCase        func(mockUseCase *mocks.FollowUserUseCase)
		expectedStatusCode int
	}{
		{
			name:               "should return 400 if request binding fails",
			requestBody:        `{"user_id": ""}`,
			mockUseCase:        func(mockUseCase *mocks.FollowUserUseCase) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "should return 500 if use case returns an internal error",
			requestBody: `{"user_id": "0f089136-3f38-4757-840c-d0c954782457", "user_id_to_follow": "12345678-3f38-4757-840c-d0c954782457"}`,
			mockUseCase: func(mockUseCase *mocks.FollowUserUseCase) {
				mockUseCase.
					On("FollowUser", mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "should return 202 if follow is successful",
			requestBody: `{"user_id": "0f089136-3f38-4757-840c-d0c954782457", "user_id_to_follow": "12345678-3f38-4757-840c-d0c954782457"}`,
			mockUseCase: func(mockUseCase *mocks.FollowUserUseCase) {
				mockUseCase.
					On("FollowUser", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			expectedStatusCode: http.StatusAccepted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(mocks.FollowUserUseCase)
			tt.mockUseCase(mockUseCase)
			handler := NewFollowUserHandler(mockUseCase)

			router := gin.Default()
			router.POST("/follow", handler.Handle)

			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/follow", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(recorder, req)

			require.Equal(t, tt.expectedStatusCode, recorder.Code)
			mockUseCase.AssertExpectations(t)
		})
	}
}
