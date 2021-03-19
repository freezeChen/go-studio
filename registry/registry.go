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
	ListServices() (map[string][]string, error)

	WatchServices()
}
