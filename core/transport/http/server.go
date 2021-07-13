package http

import (
	"context"
	"github.com/freezeChen/go-studio/util/addr"
	"net/http"
)

type Server struct {
	s       *http.Server
	h       http.Handler
	address string
}

func NewServer(address string) *Server {
	var s = new(Server)
	s.address = address
	return s
}

func (s *Server) Handle(h http.Handler) {
	s.h = h
}

func (s *Server) Endpoint() (string, string, error) {
	readIp, err := addr.Extract("")
	if err != nil {
		return "", "", err
	}

	return "HTTP", readIp + s.address, err
}

func (s *Server) Start() error {
	s.s = &http.Server{
		Addr: s.address,
		Handler: s.h,
	}
	return s.s.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.s.Shutdown(context.Background())
}
