package event

import (
	"reflect"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/while-loop/levit/common/test"
	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
	"github.com/while-loop/levit/hub/stream"
)

func TestMessage(t *testing.T) {
	msg := &proto.HubMessage{
		Event: &proto.HubMessage_EventMessage{
			EventMessage: &proto.EventMessage{
				Message: "hello",
			},
		},
	}

	h := hub.New(nil)
	go h.Start()
	m := stream.NewMock()
	c := hub.NewConn(h, m)
	h.Register(c)

	GetHandler()[reflect.TypeOf(&proto.HubMessage_EventMessage{})](c, msg)

	assert.True(t, test.RanWithinTimeout(100*time.Millisecond, func() {
		assert.Equal(t, msg, <-m.SendBuf)
	}))
}
