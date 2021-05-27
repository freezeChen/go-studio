package main

import (
	"context"
	proxy2 "go-studio/core/proxy"
	registry2 "go-studio/core/registry"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	etcdRegistry, err := registry2.NewEtcdRegistry(ctx, []string{"localhost:2379"})
	if err != nil {
		panic(err)
	}
	services, err := etcdRegistry.ListServices()
	if err != nil {
		panic(err)
	}
	for k, v := range services {
		log.Printf("%s ,%v", k, v)
	}
	//parse, err := url.Parse("http://localhost:8081")


	http.ListenAndServe(":8082", proxy2.NewProxy(etcdRegistry))
	//http.ListenAndServe(":8082", httputil.NewSingleHostReverseProxy(parse))
}
