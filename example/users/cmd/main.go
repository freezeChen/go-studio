package main

import (
	studio "github.com/freezeChen/go-studio"
	"github.com/freezeChen/go-studio/example/users/internal/interfaces"
)

func main() {
	//ctx := context.Background()
	//etcdRegistry, err := registry.NewEtcdRegistry(ctx, []string{"localhost:2379"})
	//if err != nil {
	//	panic(err)
	//}

	server := interfaces.NewHttpServer()

	app := studio.New(
		studio.Name("users"),
		studio.Version("1.0"),
		//studio.Registrar(etcdRegistry),
		studio.Server(
			server,
		),
	)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
