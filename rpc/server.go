package rpc

import (
	"context"
	"go-studio/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server *grpc.Server
	port   string
}

func NewServer(port string) *Server {
	s := &Server{}
	s.server = grpc.NewServer()
	return s
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	err = s.server.Serve(listen)
	return err
}

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
