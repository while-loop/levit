package registry

import (
	"github.com/while-loop/levit/common/log"
)

type stubReg struct {
	services map[string]Service
}

func NewStub() Registry {
	log.Info("Running stub registry")
	return &stubReg{services: map[string]Service{}}
}

func (s *stubReg) Register(srv Service) error {
	srvKey := key(srv)

	if _, exists := s.services[srvKey]; !exists {
		s.services[srvKey] = srv
	}

	for uuid, inst := range srv.Instances {
		s.services[srvKey].Instances[uuid] = inst
	}

	return nil
}

func (s *stubReg) Deregister(srv Service) error {
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

func (s *stubReg) GetServices() ([]Service, error) {
	srvs := make([]Service, 0)
	for _, srv := range s.services {
		srvs = append(srvs, srv)
	}

	return srvs, nil
}

func (s *stubReg) GetService(serviceName, version string) (Service, error) {
	err := ErrServiceDNE
	for _, srv := range s.services {
		if srv.Name == serviceName && srv.Version == version {
			return srv, nil
		}
	}

	return Service{}, err
}

func (s *stubReg) Name() string {
	return "stub"
}

func key(service Service) string {
	return service.Name + "-" + service.Version
}
