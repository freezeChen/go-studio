package go_studio

import (
	"context"
	registry2 "go-studio/core/registry"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	ctx := context.Background()
	etcdRegistry, err := registry2.NewEtcdRegistry(ctx, []string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	app := New(Name("users"),
		Version("1.0"),
		Registrar(etcdRegistry),

	)

	time.AfterFunc(2*time.Second, func() {
		app.Stop()
	})

	if err := app.Start(); err != nil {
		t.Fatal(err)
	}
}
