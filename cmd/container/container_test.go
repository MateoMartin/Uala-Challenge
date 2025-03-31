package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadContainer(t *testing.T) {
	container := LoadContainer()

	assert.NotNil(t, container.GetStatusHandler, "GetStatusHandler should not be nil")
}
