package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRanTimeouts(t *testing.T) {
	assert.True(t, RanWithinTimeout(50*time.Millisecond, func() {
		time.Sleep(25 * time.Millisecond)
	}))
	assert.False(t, RanWithinTimeout(50*time.Millisecond, func() {
		time.Sleep(75 * time.Millisecond)
	}))
}

func TestEqualsWithinTimeout(t *testing.T) {
	assert.False(t, EqualsWithinTimeout(50*time.Millisecond, func() bool {
		time.Sleep(25 * time.Millisecond)
		return false
	}))
	assert.True(t, EqualsWithinTimeout(50*time.Millisecond, func() bool {
		time.Sleep(25 * time.Millisecond)
		return true
	}))

	assert.False(t, EqualsWithinTimeout(50*time.Millisecond, func() bool {
		time.Sleep(75 * time.Millisecond)
		return true
	}))

	assert.False(t, EqualsWithinTimeout(50*time.Millisecond, func() bool {
		time.Sleep(75 * time.Millisecond)
		return false
	}))
}
