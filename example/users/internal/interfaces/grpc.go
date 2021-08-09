package interfaces

import (
	"context"
	"github.com/freezeChen/go-studio/example/proto"
	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	var opts = []grpc.ServerOption{}

	opts = append(opts, grpc.Netwo)

	server := grpc.NewServer()

	proto.RegisterGreetServer(server)
}

type s struct {
}

func (s s) Morning(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	panic("implement me")
}

func (s s) Night(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	panic("implement me")
}
