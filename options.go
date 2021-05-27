package go_studio

import (
	"context"
	registry2 "go-studio/core/registry"
	transport2 "go-studio/core/transport"
)

type Option func(o *Options)

type Options struct {
	ctx       context.Context
	name      string
	version   string
	metadata  map[string]string
	registrar registry2.Registrar
	servers   []transport2.Server
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

func Server(server ...transport2.Server) Option {
	return func(o *Options) {
		o.servers = server
	}
}

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func Registrar(r registry2.Registrar) Option {
	return func(o *Options) {
		o.registrar = r
	}
}
