package go_studio

import (
	"context"
	"go-studio/registry"
	"go-studio/transport"
)

type Option func(o *Options)

type Options struct {
	ctx       context.Context
	name      string
	version   string
	metadata  map[string]string
	registrar registry.Registrar
	servers    []transport.Server
}

func newOptions(opts ...Option) Options {
	options := Options{
		ctx:       context.Background(),
		registrar: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}
	return options
}

func Name(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

func Version(version string) Option {
	return func(o *Options) {
		o.version = version
	}
}

func Server(server ...transport.Server) Option {
	return func(o *Options) {
		o.servers = server
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func Registrar(r registry.Registrar) Option {
	return func(o *Options) {
		o.registrar = r
	}
}
