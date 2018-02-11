package hub

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/while-loop/levit/common/test"
	"github.com/while-loop/levit/hub/proto"
	"github.com/while-loop/levit/hub/stream"
)

func TestHub_Register_Deregister(t *testing.T) {
	h := New(nil)
	go h.Start()

	c := NewConn(h, stream.NewMock())
	h.Register(c)

	var wg sync.WaitGroup
	wg.Add(1)
	h.connChan <- func(conns map[uint64]*Conn) {
		assert.Len(t, conns, 1)
		assert.Contains(t, conns, c.UserId)
		wg.Done()
	}

	assert.True(t, test.RanWithinTimeout(25*time.Millisecond, func() {
		wg.Wait()
	}))

	h.Deregister(c)
	wg.Add(1)
	h.connChan <- func(conns map[uint64]*Conn) {
		assert.Len(t, conns, 0)
		wg.Done()
	}

	assert.True(t, test.RanWithinTimeout(25*time.Millisecond, func() {
		wg.Wait()
	}))

	h.Stop()
}

func TestMessagesBroadcast(t *testing.T) {
	h := New(nil)
	go h.Start()

	c1 := NewConn(h, stream.NewMock())
	c2 := NewConn(h, stream.NewMock())
	h.Register(c1)
	h.Register(c2)

	var wg sync.WaitGroup
	wg.Add(1)
	h.connChan <- func(conns map[uint64]*Conn) {
		require.Len(t, conns, 2)
		wg.Done()
	}

	assert.True(t, test.RanWithinTimeout(25*time.Millisecond, func() {
		wg.Wait()
	}))

	h.Broadcast(&proto.HubMessage{
		Data: "hi",
	})

	assert.True(t, test.RanWithinTimeout(250*time.Millisecond, func() {
		m := <-c1.s.(*stream.MockStream).SendBuf
		m2 := <-c2.s.(*stream.MockStream).SendBuf
		assert.Equal(t, "hi", m.Data)
		assert.Equal(t, m, m2)
	}))
}
