package event

import (
	"reflect"

	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
)

func init() {
	handler[reflect.TypeOf(&proto.HubMessage_EventMessage{})] = eventmessage
}

func eventmessage(conn *hub.Conn, message *proto.HubMessage) {

}
