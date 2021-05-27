package rpc

import (
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

func (s *Server) Server() *grpc.Server {
	return s.server
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}
	err = s.server.Serve(listen)
	return err
}

