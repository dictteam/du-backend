package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
)

type App struct {
	container container.Container
}

func (app *App) Start() error {
	var fiberInstance *fiber.App
	container.MustResolve(app.container, &fiberInstance)

	return fiberInstance.Listen(":3000")
}

// AddController adds a controller to the app initializing its dependencies
func (app *App) AddController(controller Controller) {
	controller.Init(app.container)
}

// ShutDown gracefully shutdowns the app
func (app *App) ShutDown() error {
	return nil
}

func NewApp(container container.Container) *App {
	app := &App{
		container,
	}

	return app
}
