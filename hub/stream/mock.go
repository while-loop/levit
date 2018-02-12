package stream

import (
	"fmt"

	"math/rand"
	"time"

	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/hub/proto"
)

var (
	r = rand.New(rand.NewSource(time.Now().Unix()))
)

type MockStream struct {
	RecvBuf chan *proto.HubMessage
	SendBuf []*proto.HubMessage
	RecvErr error
	SendErr error
	Id      uint64
}

func (m *MockStream) Recv() (*proto.HubMessage, error) {
	log.Debug("MockStream: ", m.Id, " Recv: ")
	for {
		select {
		case msg := <-m.RecvBuf:
			return msg, nil
		default:
			if m.RecvErr != nil {
				return nil, m.RecvErr
			}
		}
	}
}

func (m *MockStream) Send(msg *proto.HubMessage) error {
	log.Debug("MockStream: ", m.Id, " Send: ", msg)
	m.SendBuf = append(m.SendBuf, msg)
	return m.SendErr
}

func (m *MockStream) Close() error {
	m.RecvErr = fmt.Errorf("done")
	return nil
}

func NewMock() *MockStream {
	return &MockStream{
		Id:      r.Uint64(),
		RecvBuf: make(chan *proto.HubMessage, 100),
		SendBuf: make([]*proto.HubMessage, 0),
	}
}
