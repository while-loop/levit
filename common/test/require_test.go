package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasInstalled(t *testing.T) {
	assert.True(t, HasInstalled("sh"))
	assert.False(t, HasInstalled("shasdasdasd"))
	assert.Equal(t, HasInstalled("docker"), HasDockerInstalled())
}
