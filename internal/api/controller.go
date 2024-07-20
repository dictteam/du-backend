package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golobby/container/v3"
)

type Controller interface {
	// Init Initialize controller dependencies
	Init(router fiber.Router, cont container.Container) error
	// ShutDown gracefully shutdown controller and current transactions
	ShutDown() error
}
