package stub

import (
	"github.com/while-loop/levit/liblevit/registry"
	"fmt"
)

type stubReg struct {
	services map[string]registry.Service
}

func New() registry.Registry {
	return &stubReg{services: map[string]registry.Service{}}
}

func (s *stubReg) Register(srv registry.Service) error {
	srvKey := key(srv)

	if _, exists := s.services[srvKey]; !exists {
		s.services[srvKey] = srv
	}

	for uuid, inst := range srv.Instances {
		s.services[srvKey].Instances[uuid] = inst
	}

	return nil
}

func (s *stubReg) Deregister(srv registry.Service) error {
	srvKey := key(srv)
	if _, exists := s.services[srvKey]; !exists {
		return nil
	}

	for uuid := range srv.Instances {
		delete(s.services[srvKey].Instances, uuid)
	}

	if len(s.services[srvKey].Instances) == 0 {
		delete(s.services, srvKey)
	}

	return nil
}

func (s *stubReg) GetServices() ([]registry.Service, error) {
	srvs := make([]registry.Service, 0)
	for _, srv := range s.services {
		srvs = append(srvs, srv)
	}

	return srvs, nil
}

func (s *stubReg) GetService(serviceName, version string) (registry.Service, error) {
	err := fmt.Errorf("dne")
	for _, srv := range s.services {
		if srv.Name == serviceName && srv.Version == version {
			return srv, nil
		}
	}

	return registry.Service{}, err
}

func (s *stubReg) Name() string {
	return "stub"
}

func key(service registry.Service) string {
	return service.Name + "-" + service.Version
}
