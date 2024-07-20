package users

import (
	fiber "github.com/gofiber/fiber/v2"
	container "github.com/golobby/container/v3"
)

type UserController struct {
}

// Init implements api.Controller.
func (c *UserController) Init(router fiber.Router, cont container.Container) error {
	router.Get("/", c.GetUsers)
	return nil
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"msg": "Method not implemented",
	})
}

// ShutDown implements api.Controller.
func (*UserController) ShutDown() error {
	return nil
}

func NewUsersController() *UserController {
	return &UserController{}
}
