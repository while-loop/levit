package event

import (
	"reflect"

	"github.com/while-loop/levit/hub"
)

var handler = hub.Handler{}

func GetHandler() hub.Handler {
	return handler
}

func RegisterEvent(event interface{}, eventFunc hub.EventFunc) {
	handler[reflect.TypeOf(event)] = eventFunc
}
