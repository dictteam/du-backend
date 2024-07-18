package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"msg": "Hello, World!",
		})
	})

	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong!")
	})

	log.Fatal(app.Listen(":3000"))
}
