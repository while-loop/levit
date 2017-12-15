package registry

type Registry interface {
	Register(service Service) error
	Deregister(service Service) error
	GetServices() ([]Service, error)
	GetService(serviceName, version string) (Service, error)
	Name() string
}

type Service struct {
	Name      string
	Version   string
	Instances map[string]Instance
}

type Instance struct {
	UUID string
	IP   string
	Port int
}

var (
	registry Registry = nil
)

func Use(rgstry Registry) {
	registry = rgstry
}

func Register(srv Service) error                             { return registry.Register(srv) }
func Deregister(srv Service) error                           { return registry.Deregister(srv) }
func GetServices() ([]Service, error)                         { return registry.GetServices() }
func GetService(serviceName, version string) (Service, error) { return registry.GetService(serviceName, version) }
func Name() string                                            { return registry.Name() }