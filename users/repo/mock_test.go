package repo

import "testing"

func TestMockImpl(t *testing.T) {
	testImpl(t, NewMockRepo())
}
