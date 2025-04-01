package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"uala-challenge/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func mockHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func TestSetupRouter(t *testing.T) {
	router := SetupRouter(config.LoadConfig(), mockHandler, mockHandler)

	tests := []struct {
		method   string
		endpoint string
		status   int
	}{
		{"GET", "/status", http.StatusOK},
		{"POST", "/tweet", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.method+" "+tt.endpoint, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.endpoint, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}
