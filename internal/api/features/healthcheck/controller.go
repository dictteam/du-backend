package healthcheck

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	container "github.com/golobby/container/v3"
)

type HealthCheckController struct {
}

// Init implements api.Controller.
func (*HealthCheckController) Init(cont container.Container) error {
	var router fiber.Router
	err := cont.Resolve(&router)
	if err != nil {
		return nil
	}
	group := router.Group("/healthcheck")
	group.Get("/", func(ctx *fiber.Ctx) error {
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
