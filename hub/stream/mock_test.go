package stream

import (
	"testing"

	"time"

	"github.com/stretchr/testify/require"
	"github.com/while-loop/levit/hub/proto"
)

func TestMockSend(t *testing.T) {
	a := require.New(t)
	m := NewMock()
	ms := []*proto.HubMessage{{Uid: 5}, {Uid: 6}}

	time.AfterFunc(1000*time.Millisecond, func() {
		a.FailNow("timeout reached")
	})

	for _, msg := range ms {
		m.Send(msg)
	}

	for idx := range ms {
		a.Equal(ms[idx], <-m.SendBuf)
	}
}

func TestMockRecv(t *testing.T) {
	a := require.New(t)
	m := NewMock()

	time.AfterFunc(1000*time.Millisecond, func() {
		a.FailNow("timeout reached")
	})

	ms := []*proto.HubMessage{{Uid: 5}, {Uid: 6}}
	for _, msg := range ms {
		m.RecvBuf <- msg
	}

	for idx := range ms {
		msg, err := m.Recv()
		a.NoError(err)
		a.Equal(ms[idx], msg)
	}
}
