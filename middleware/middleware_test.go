package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"uala-challenge/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestMethodNotAllowed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.HandleMethodNotAllowed = true
	router.NoMethod(MethodNotAllowed)
	router.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	req := httptest.NewRequest("POST", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestContentTypeValidator(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ContentTypeValidator("application/json"))

	router.POST("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	t.Run("should return 400 if Content-Type is not application/json", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "text/plain")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error": "Bad Request - Content-Type must be application/json"}`, w.Body.String())
	})

	t.Run("should pass if Content-Type is application/json", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestLoggerMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	observedZapCore, observedLogs := observer.New(zapcore.InfoLevel)
	zapLogger := zap.New(observedZapCore).Sugar()
	logger.Log = zapLogger

	router := gin.New()
	router.Use(LoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	t.Run("should log request details", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		router.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		logs := observedLogs.All()
		require.Len(t, logs, 1)

		entry := logs[0]
		assert.Equal(t, "new_request", entry.Message)
		assert.Equal(t, zapcore.InfoLevel, entry.Level)

	})
}
