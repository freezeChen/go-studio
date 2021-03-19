package main

import (
	"context"
	go_studio "go-studio"
	"go-studio/example/users/internal/interfaces"
	"go-studio/registry"
)

func main() {
	ctx := context.Background()
	etcdRegistry, err := registry.NewEtcdRegistry(ctx, []string{"127.0.0.1:2379"})
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
