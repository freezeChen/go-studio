package registry

type Service struct {
	Name     string
	Port     string
	MetaData map[string]string
}

type Registry interface {
	RegistryService(Service) error
	//UnRegistryService(Service) error
	ListServices()(map[string][]string, error)
	WatchServices()
}
