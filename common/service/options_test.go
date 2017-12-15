package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLaddr(t *testing.T) {
	o := Options{
		IP:   "localhost",
		Port: 5353,
	}
	require.Equal(t, "localhost:5353", o.laddr())
}
