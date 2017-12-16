package registry

import (
	"fmt"
	"time"
)

type Registry interface {
	Register(service Service) error
	Deregister(service Service) error
	GetServices() ([]Service, error)
	GetService(serviceName, version string) ([]Service, error)
	Name() string
}

type Service struct {
	Name    string
	UUID    string
	Version string
	IP      string
	Port    int
	TTL     time.Duration
}

func (s Service) Key() string {
	return fmt.Sprintf("%s-%s-%s", s.Name, s.Version, s.UUID)
}

var (
	registry      Registry = nil
	ErrServiceDNE          = fmt.Errorf("service does not exist")
)

func Use(rgstry Registry) {
	registry = rgstry
}

func Register(srv Service) error      { return registry.Register(srv) }
func Deregister(srv Service) error    { return registry.Deregister(srv) }
func GetServices() ([]Service, error) { return registry.GetServices() }
func GetService(serviceName, version string) ([]Service, error) {
	return registry.GetService(serviceName, version)
}
func Name() string { return registry.Name() }
