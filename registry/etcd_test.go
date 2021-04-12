package registry

import (
	"context"
	"testing"
)

func TestNewEtcdRegistry(t *testing.T) {
	registry, err := NewEtcdRegistry(context.TODO(), []string{"localhost:2379"})
	if err != nil {
		panic(err)
	}

	registry.RegistryService(&Service{Id: "123333"})

	services, err := registry.ListServices()
	if err != nil {
		panic(err)
	}
	t.Log(services)
}