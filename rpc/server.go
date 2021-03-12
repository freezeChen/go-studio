package rpc

import (
	"context"
	"go-studio/proto"
	"google.golang.org/grpc"
)

func a() {
	server := grpc.NewServer()
	proto.RegisterGreetServer(server, &Greet{})

server.Serve("")
}

type Greet struct {
}

func (g *Greet) Morning(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	panic("implement me")
}

func (g *Greet) Night(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	panic("implement me")
}
