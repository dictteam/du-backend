package healthcheck

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	container "github.com/golobby/container/v3"
)

type HealthCheckController struct {
}

// Init implements api.Controller.
func (*HealthCheckController) Init(router fiber.Router, cont container.Container) error {
	router.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).SendString("ok")
	})

	return nil
}

// ShutDown implements api.Controller.
func (*HealthCheckController) ShutDown() error {
	return nil
}

func NewController() *HealthCheckController {
	return &HealthCheckController{}
}
