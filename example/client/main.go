package main

import (
	"context"
	"fmt"
	go_studio "go-studio"
	"go-studio/registry"
	"time"
)

func main() {
	ctx := context.Background()
	etcdRegistry, err := registry.NewEtcdRegistry(ctx, []string{"localhost:2379"})
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(2 * time.Second)
		services, err := etcdRegistry.ListServices()
		if err != nil {
			panic(err)
		}
		for k, strings := range services {
			fmt.Println(k, strings)
		}
	}()

	app := go_studio.New(
		go_studio.Name("1users"),
		go_studio.Version("1.0"),
		go_studio.Registrar(etcdRegistry),
	)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
