package main

import (
	"log"

	"github.com/4strodev/du_backend/internal/api"
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

		app.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.JSON(fiber.Map{
				"msg": "Hello, world!",
			})
		})

		app.Get("/ping", func(ctx *fiber.Ctx) error {
			return ctx.SendString("pong!")
		})

		return app
	})

	// Here dependencies will be attached to the container
	app := api.NewApp(diContainer)

	log.Fatal(app.Start())
}
