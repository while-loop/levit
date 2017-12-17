package event

import (
	"github.com/while-loop/levit/hub"
)

var router = hub.Router{}

func GetRouter() hub.Router {
	return router
}
