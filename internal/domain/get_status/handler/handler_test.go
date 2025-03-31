package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetStatusHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/status", NewGetStatusHandler().Handle)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/status", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
