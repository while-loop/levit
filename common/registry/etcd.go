package registry

type EtcdRegistry struct {
}

// def implementation is kubernetes
// TODO switch def impl from Stub to kubernetes
func NewK8s() Registry {
	return &EtcdRegistry{}
}

func (k *EtcdRegistry) Register(service Service) error {
	panic("implement me")
}

func (k *EtcdRegistry) Deregister(service Service) error {
	panic("implement me")
}

func (k *EtcdRegistry) GetServices() ([]Service, error) {
	panic("implement me")
}

func (k *EtcdRegistry) GetService(serviceName, version string) (Service, error) {
	panic("implement me")
}

func (k *EtcdRegistry) Name() string {
	panic("implement me")
}
