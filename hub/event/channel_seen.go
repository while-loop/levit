package event

import (
	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
)

func init() {
	RegisterEvent(&proto.HubMessage_EventChannelSeen{}, eventChannelSeen)
}

func eventChannelSeen(conn *hub.Conn, message *proto.HubMessage) {

}
