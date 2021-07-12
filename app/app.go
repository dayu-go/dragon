package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/registry"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// App is an application components lifecycle manager
type App struct {
	opts     options
	ctx      context.Context
	cancel   func()
	instance *registry.ServiceInstance
}

// New create an application lifecycle manager.
func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		logger: log.NewHelper(log.DefaultLogger),
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	if id, err := uuid.NewUUID(); err == nil {
		options.id = id.String()
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		opts:   options,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	if len(a.opts.endpoints) == 0 {
		for _, srv := range a.opts.servers {
			e, err := srv.Endpoint()
			if err == nil {
				return nil, err
			}
			a.opts.endpoints = append(a.opts.endpoints, e)
		}
	}
	return &registry.ServiceInstance{
		ID:        a.opts.id,
		Name:      a.opts.name,
		Version:   a.opts.version,
		Metadata:  a.opts.metadata,
		Endpoints: a.opts.endpoints,
	}, nil
}

func (a *App) Run() error {
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}
	a.opts.logger.Log(log.LevelInfo,
		"service_id", a.opts.id,
		"service_name", a.opts.name,
		"version", a.opts.version,
		"metadata", a.opts.metadata,
		"endpoints", a.opts.endpoints,
	)
	g, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	for _, srv := range a.opts.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop(ctx)
		})
		wg.Add(1)
		g.Go(func() error {
			err = srv.Start(ctx)
			wg.Done()
			return err
		})
	}
	wg.Wait()
	if a.opts.registrar != nil {
		if err := a.opts.registrar.Register(a.opts.ctx, instance); err != nil {
			return err
		}
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.sigs...)
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
	if a.opts.registrar != nil {
		if err := a.opts.registrar.Deregister(a.opts.ctx, a.instance); err != nil {
			return err
		}
	}
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
