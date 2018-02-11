package event

import (
	"github.com/while-loop/levit/hub"
)

var handler = hub.Handler{}

func GetHandler() hub.Handler {
	return handler
}
