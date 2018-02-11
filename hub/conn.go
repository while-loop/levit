package hub

import (
	"github.com/while-loop/levit/hub/proto"

	"math/rand"
	"time"

	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/hub/stream"
)

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

type Conn struct {
	UserId uint64
	Parent *Hub
	s      stream.Stream
}

func NewConn(hub *Hub, s stream.Stream) *Conn {
	c := &Conn{
		UserId: r.Uint64(),
		Parent: hub,
		s:      s,
	}

	return c
}

func (c *Conn) Send(message *proto.HubMessage) error {
	return c.s.Send(message)
}

func (c *Conn) Loop() {
	c.readLoop()
	c.Close()
}

func (c *Conn) readLoop() {
	defer c.s.Close()
	var msg *proto.HubMessage
	var err error
	for {
		msg, err = c.s.Recv()
		if err != nil {
			log.Errorf("failed read from conn stream %v", err)
			return
		}

		// TODO do router message to handler
		log.Debug(msg)
	}
}

func (c *Conn) Close() {
	c.s.Close()
}

func (c *Conn) Contains(message *proto.HubMessage) bool {
	return true
}
