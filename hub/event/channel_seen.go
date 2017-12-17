package event

import (
	"reflect"

	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
)

func init() {
	router[reflect.TypeOf(&proto.HubMessage_EventChannelSeen{})] = eventChannelSeen
}

func eventChannelSeen(conn *hub.Conn, message *proto.HubMessage) {

}
