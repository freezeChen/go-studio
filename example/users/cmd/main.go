package main

import (
	"context"
	go_studio "go-studio"
	registry2 "go-studio/core/registry"
	"go-studio/example/users/internal/interfaces"
)

func main() {
	ctx := context.Background()
	etcdRegistry, err := registry2.NewEtcdRegistry(ctx, []string{"localhost:2379"})
	if err != nil {
		panic(err)
	}

	server := interfaces.NewHttpServer()

	app := go_studio.New(
		go_studio.Name("users"),
		go_studio.Version("1.0"),
		go_studio.Registrar(etcdRegistry),
		go_studio.Server(
			server,
		),
	)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
