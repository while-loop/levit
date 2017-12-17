package stream

import (
	"io"

	"github.com/while-loop/levit/hub/proto"
)

type Stream interface {
	Recv() (*proto.HubMessage, error)
	Send(msg *proto.HubMessage) error
	io.Closer
}
