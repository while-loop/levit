package hub

import (
	"fmt"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/require"
	"github.com/while-loop/levit/hub/proto"
	"github.com/while-loop/levit/hub/stream"
)

func TestCloseStopsLoops(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)

	mock := stream.NewMock().(*stream.MockStream)
	c := NewConn(nil, mock)
	time.AfterFunc(100*time.Millisecond, func() {
		mock.Close()
		time.AfterFunc(500*time.Millisecond, func() {
			a.FailNow("failed to get msg")
		})
	})
	c.Loop()
}

func TestConnSendsToChannel(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)
	mock := stream.NewMock().(*stream.MockStream)
	c := NewConn(nil, mock)
	msg := &proto.HubMessage{}
	c.Send() <- msg

	// make sure message doesnt send before starting Loop()
	select {
	case <-mock.SendBuf:
		t.FailNow()
	case <-time.After(100 * time.Millisecond):
	}

	time.AfterFunc(100*time.Millisecond, func() {
		time.AfterFunc(500*time.Millisecond, func() {
			mock.Close()
			a.FailNow("failed to get msg")
		})
		a.Equal(msg, <-mock.SendBuf)
		mock.Close()
	})
	c.Loop()
}

func TestConnErrorInSendEndsConn(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)
	mock := stream.NewMock().(*stream.MockStream)
	c := NewConn(nil, mock)

	time.AfterFunc(100*time.Millisecond, func() {
		mock.SendErr = fmt.Errorf("send error")
		c.Send() <- nil
		time.AfterFunc(1000*time.Millisecond, func() {
			mock.Close()
			a.FailNow("failed to get msg")
		})
	})
	c.Loop()

}

func TestConnRecvs(t *testing.T) {
	defer leaktest.Check(t)()
	a := require.New(t)
	mock := stream.NewMock().(*stream.MockStream)
	c := NewConn(nil, mock)
	msg := &proto.HubMessage{}

	time.AfterFunc(100*time.Millisecond, func() {
		mock.RecvBuf <- msg
		time.AfterFunc(1000*time.Millisecond, func() {
			mock.Close()
			a.FailNow("failed to get msg")
		})
		mock.Close()
	})
	c.Loop()

}
