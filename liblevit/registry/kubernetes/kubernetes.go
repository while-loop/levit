package kubernetes

import (
	"github.com/while-loop/levit/liblevit/registry"
)


// def implementation is kubernetes
func New() registry.Registry {
	return registry.New()
}
