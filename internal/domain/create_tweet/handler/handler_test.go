package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"uala-challenge/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"uala-challenge/internal/domain/create_tweet/handler/mocks"
)

func TestCreateTweetHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name               string
		requestBody        string
		mockUseCase        func(mockUseCase *mocks.CreateTweetUseCase)
		expectedStatusCode int
	}{
		{
			name:               "should return 400 if request binding fails",
			requestBody:        `{"content": ""}`,
			mockUseCase:        func(mockUseCase *mocks.CreateTweetUseCase) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "should return 400 if tweet is invalid",
			requestBody:        fmt.Sprintf(`{"content": "%s", "user_id": "0f089136-3f38-4757-840c-d0c954782457"}`, strings.Repeat("a", model.MaxCharacters+1)),
			mockUseCase:        func(mockUseCase *mocks.CreateTweetUseCase) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "should return 500 if use case returns an internal error",
			requestBody: `{"content": "Hello World", "user_id": "0f089136-3f38-4757-840c-d0c954782457"}`,
			mockUseCase: func(mockUseCase *mocks.CreateTweetUseCase) {
				mockUseCase.
					On("CreateTweet", mock.Anything, mock.Anything).
					Return(errors.New("internal error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "should return 201 if tweet creation is successful",
			requestBody: `{"content": "Hello World", "user_id": "0f089136-3f38-4757-840c-d0c954782457"}`,
			mockUseCase: func(mockUseCase *mocks.CreateTweetUseCase) {
				mockUseCase.
					On("CreateTweet", mock.Anything, mock.Anything).
					Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUseCase := new(mocks.CreateTweetUseCase)
			tt.mockUseCase(mockUseCase)
			handler := NewCreateTweetHandler(mockUseCase)

			router := gin.Default()
			router.POST("/tweets", handler.Handle)

			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/tweets", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(recorder, req)

			require.Equal(t, tt.expectedStatusCode, recorder.Code)
			mockUseCase.AssertExpectations(t)
		})
	}
}
