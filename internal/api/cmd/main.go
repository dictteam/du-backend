package main

import (
	"log"

	"github.com/4strodev/du_backend/internal/api"
	"github.com/4strodev/du_backend/internal/api/features/healthcheck"
	"github.com/4strodev/du_backend/internal/api/features/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/golobby/container/v3"
)

func main() {
	diContainer := container.New()

	diContainer.Singleton(func() *fiber.App {
		app := fiber.New()

		app.Use(recover.New())
		app.Use(logger.New())

		return app
	})

	// Here dependencies will be attached to the container
	app := api.NewApp(diContainer)

	app.AddController("", healthcheck.NewController())
	app.AddController("users", users.NewUsersController())

	log.Fatal(app.Start())
}
