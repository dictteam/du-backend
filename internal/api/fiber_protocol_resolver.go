package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golobby/container/v3"
)

type FiberProtocolResolver struct {
	app *fiber.App
}

// Resolve implements ProtocolAdapter.
func (f *FiberProtocolResolver) PrepareDependency(cont container.Container) error {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	f.app = app

	cont.Singleton(func() fiber.Router {
		return app
	})

	return nil
}

// Start implements ProtocolAdapter.
func (f *FiberProtocolResolver) Start() error {
	return f.app.Listen(":3000")
}
