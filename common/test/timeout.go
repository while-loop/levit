package test

import (
	"time"
)

func RanWithinTimeout(d time.Duration, fn func()) bool {
	done := make(chan struct{})
	go func() {
		fn()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return true
	case <-time.After(d):
		return false
	}
}

func EqualsWithinTimeout(d time.Duration, fn func() bool) bool {
	return RanWithinTimeout(d, func() {
		for !fn() {
			time.Sleep(25 * time.Millisecond)
		}
	})
}
