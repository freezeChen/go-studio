package registry

type Service struct {
	Id        string
	Name      string
	Version   string
	MetaData  map[string]string
	Endpoints map[string]string
}

type Registrar interface {
	RegistryService(*Service) error
	UnRegistryService(*Service) error
	GetService(serviceName string) ([]*Service, error)
	ListServices() (map[string][]*Service, error)
	WatchServices()
}
