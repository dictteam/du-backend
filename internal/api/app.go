package api

import (
	"errors"
	"fmt"
	"github.com/golobby/container/v3"
)

type App struct {
	serviceContainer container.Container
	controllers      []Controller
	protocolAdapters []ProtocolAdapter
}

func (app *App) Start() error {
	errChan := make(chan error)

	for _, adapter := range app.protocolAdapters {
		go func(errChan chan error, adapter ProtocolAdapter) {
			err := adapter.Start()
			if err != nil {
				errChan <- err
			}
		}(errChan, adapter)
	}

	return <-errChan
}

func (app *App) AddProtocolAdapter(adapter ProtocolAdapter) error {
	err := adapter.PrepareDependencies(app.serviceContainer)
	if err != nil {
		return fmt.Errorf("error perparing adapter: %w", err)
	}

	app.protocolAdapters = append(app.protocolAdapters, adapter)

	return nil
}

// AddController adds a controller to the app initializing its dependencies
func (app *App) AddController(controller Controller) error {
	err := controller.Init(app.serviceContainer)
	if err != nil {
		return fmt.Errorf("error initializing controllers: %w", err)
	}

	app.controllers = append(app.controllers, controller)

	return nil
}

// ShutDown gracefully shutdowns the app
func (app *App) ShutDown() error {
	var errs []error = make([]error, 0)
	for _, controller := range app.controllers {
		err := controller.ShutDown()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func NewApp(serviceContainer container.Container) *App {
	app := &App{
		serviceContainer: serviceContainer,
		controllers:      make([]Controller, 0),
		protocolAdapters: make([]ProtocolAdapter, 0),
	}

	return app
}
