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
	s.services[srv.Key()] = srv
	return nil
}

func (s *stubReg) Deregister(srv Service) error {
	delete(s.services, srv.Key())
	return nil
}

func (s *stubReg) GetServices() ([]Service, error) {
	srvs := make([]Service, 0)
	for _, srv := range s.services {
		srvs = append(srvs, srv)
	}

	return srvs, nil
}

func (s *stubReg) GetService(serviceName, version string) ([]Service, error) {
	srvcs := make([]Service, 0)
	for _, srv := range s.services {
		if srv.Name == serviceName && srv.Version == version {
			srvcs = append(srvcs, srv)
		}
	}

	if len(srvcs) <= 0 {
		return nil, ErrServiceDNE
	}

	return srvcs, nil
}

func (s *stubReg) Name() string {
	return "stub"
}
