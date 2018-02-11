package event

import (
	"reflect"

	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
)

func init() {
	handler[reflect.TypeOf(&proto.HubMessage_EventMessage{})] = eventMessage
}

func eventMessage(conn *hub.Conn, message *proto.HubMessage) {
	log.Debugf("recvd message from %d: %s", conn.UserId, message.GetEventMessage().Message)
	conn.Parent.Broadcast(message)
}
