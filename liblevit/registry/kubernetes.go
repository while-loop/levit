package registry

type kubernetesReg struct {
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

func (k *kubernetesReg) GetService(serviceName, version string) (Service, error){
	panic("implement me")
}

func (k *kubernetesReg) Name() string {
	panic("implement me")
}

func newKubernetes() Registry {
	return nil
}
