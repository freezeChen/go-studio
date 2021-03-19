package http

import (
	"context"

	"github.com/gorilla/mux"
	"go-studio/util/addr"
	"net"
	"net/http"
)

type Server struct {
	s       *http.Server
	router  *mux.Router
	address string
}

func NewServer(address string) *Server {
	var s = new(Server)
	s.address = address
	s.router = mux.NewRouter()
	return s
}

func (s *Server) Handle(path string, h http.Handler) {
	s.router.Handle(path, h)
}

func (s *Server) Endpoint() (string, string, error) {
	extract, err := addr.Extract(s.address)
	return "HTTP", extract, err
}

func (s *Server) Start() error {
	s.s = &http.Server{
		Handler: s.router,
	}
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	return s.s.Serve(l)
}

func (s *Server) Stop() error {
	return s.s.Shutdown(context.Background())
}
