package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"uala-challenge/utils/logger"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
}

func ContentTypeValidator(contentType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != contentType {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Bad Request - Content-Type must be %s", contentType)})
			c.Abort()
			return
		}
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		logger.GetLogger().Infow("new_request",
			"request_id", requestid.Get(c),
			"status", statusCode,
			"method", method,
			"path", path,
			"latency", latency,
			"error", errorMessage,
		)
	}
}

func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-ctx.Done():
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": "request timed out"})
			c.Abort()
		case <-done:
		}
	}
}
