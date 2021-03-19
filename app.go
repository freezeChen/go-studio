package go_studio

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go-studio/registry"
	"go-studio/transport"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Id      string
	ctx     context.Context
	opts    Options
	cancel  func()
	service *registry.Service
	servers []transport.Server
}

func New(opts ...Option) *App {
	id := uuid.New().String()
	options := newOptions(opts...)
	ctx, cancel := context.WithCancel(options.ctx)

	enpoints := make(map[string]string, 0)
	for _, server := range options.servers {
		if kind, enpoint, err := server.Endpoint(); err == nil {
			enpoints[kind] = enpoint
		}
	}
	service := &registry.Service{
		Id:        id,
		Name:      options.name,
		Endpoints: enpoints,
		MetaData:  options.metadata,
	}

	return &App{
		Id:      id,
		ctx:     ctx,
		opts:    options,
		cancel:  cancel,
		servers: options.servers,
		service: service,
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
	if a.opts.registrar != nil {
		if err := a.opts.registrar.RegistryService(a.service); err != nil {
			return err
		}
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

	log.Printf("service_id:%s,service_name:%s,version:%s is running", a.Id, a.opts.name, a.opts.version)

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	if a.opts.registrar != nil {
		if err := a.opts.registrar.UnRegistryService(a.service); err != nil {
			fmt.Println(fmt.Printf("unRegistryService is error(%+v)", err))
		}
	}
	if a.cancel != nil {
		a.cancel()
	}

	return nil
}
