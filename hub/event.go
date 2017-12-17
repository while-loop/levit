package hub

import (
	"reflect"

	"github.com/while-loop/levit/hub/proto"
)

type EventFunc func(conn *Conn, message *proto.HubMessage)
type Router map[reflect.Type]EventFunc
