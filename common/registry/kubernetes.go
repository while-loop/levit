package registry

type kubernetesReg struct {
}

// def implementation is kubernetes
// TODO switch def impl from Stub to kubernetes
func NewK8s() Registry {
	return &kubernetesReg{}
}

func (k *kubernetesReg) Register(service Service) error {
	panic("implement me")
}

func (k *kubernetesReg) Deregister(service Service) error {
	panic("implement me")
}

func (k *kubernetesReg) GetServices() ([]Service, error) {
	panic("implement me")
}

func (k *kubernetesReg) GetService(serviceName, version string) (Service, error) {
	panic("implement me")
}

func (k *kubernetesReg) Name() string {
	panic("implement me")
}
