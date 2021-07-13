package main

import (
	"context"
	go_studio "github.com/freezeChen/go-studio"
	registry2 "github.com/freezeChen/go-studio/core/registry"
	"github.com/freezeChen/go-studio/example/users/internal/interfaces"
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
