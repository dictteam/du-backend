package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
)

type ControllerEntry struct {
	route      string
	controller Controller
}

type App struct {
	serviceContainer container.Container
	controllers      []ControllerEntry
	httpServer       *fiber.App
}

func (app *App) Start() error {
	for _, controllerEntry := range app.controllers {
		controller := controllerEntry.controller
		router := app.httpServer.Group(controllerEntry.route)

		err := controller.Init(router, app.serviceContainer)
		if err != nil {
			return fmt.Errorf("error initializing controllers: %w", err)
		}
	}

	return app.httpServer.Listen(":3000")
}

// AddController adds a controller to the app initializing its dependencies
func (app *App) AddController(route string, controller Controller) {
	app.controllers = append(app.controllers, ControllerEntry{route, controller})
}

// ShutDown gracefully shutdowns the app
func (app *App) ShutDown() error {
	var errs []error = make([]error, 0)
	for _, controllerEntry := range app.controllers {
		controller := controllerEntry.controller
		err := controller.ShutDown()
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func NewApp(serviceContainer container.Container) *App {
	var fiberInstance *fiber.App
	container.MustResolve(serviceContainer, &fiberInstance)

	app := &App{
		serviceContainer: serviceContainer,
		controllers:      make([]ControllerEntry, 0),
		httpServer:       fiberInstance,
	}

	return app
}
