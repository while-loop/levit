package hub

import (
	"github.com/while-loop/levit/hub/proto"

	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/hub/stream"
)

type Conn struct {
	UserId uint64
	in     chan *proto.HubMessage
	Parent *Hub
	s      stream.Stream
}

func NewConn(hub *Hub, s stream.Stream) *Conn {
	c := &Conn{
		UserId: 0,
		in:     make(chan *proto.HubMessage, DefaultBufferedChannelSize),
		Parent: hub,
		s:      s,
	}

	return c
}

func (c *Conn) Send() chan<- *proto.HubMessage {
	return c.in
}

func (c *Conn) Loop() {
	go c.writeLoop()
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

func (c *Conn) writeLoop() {
	defer c.s.Close()

	for msg := range c.in {
		err := c.s.Send(msg)
		if err != nil {
			log.Error("failed to send payload to client")
			break
		}
	}
}

func (c *Conn) Close() {
	close(c.in)
}

func (c *Conn) Contains(message interface{}) bool {
	return true
}
