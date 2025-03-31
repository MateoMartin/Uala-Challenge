package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	Log = nil

	t.Run("should initialize the logger if not already set", func(t *testing.T) {
		logger := GetLogger()
		assert.NotNil(t, logger)
	})

	t.Run("should return the same logger instance if already set", func(t *testing.T) {
		firstLogger := GetLogger()
		secondLogger := GetLogger()
		assert.Equal(t, firstLogger, secondLogger)
	})
}
