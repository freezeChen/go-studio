package go_studio

import (
	"context"
	"errors"
	"go-studio/registry"
	"go-studio/transport"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	ctx     context.Context
	opts    Options
	cancel  func()
	service *registry.Service
	servers []transport.Server
}

func New(opts ...Option) *App {
	options := newOptions(opts...)
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		ctx:     ctx,
		opts:    options,
		cancel:  cancel,
		servers: options.servers,
		service: nil,
	}
}

func (a *App) Start() error {
	g, ctx := errgroup.WithContext(a.ctx)

	for _, srv := range a.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}...)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Stop() error {

	if a.cancel != nil {
		a.cancel()
	}

	return nil
}
