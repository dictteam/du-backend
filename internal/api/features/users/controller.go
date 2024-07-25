package users

import (
	fiber "github.com/gofiber/fiber/v2"
	container "github.com/golobby/container/v3"
)

type UserController struct {
}

// Init implements api.Controller.
func (c *UserController) Init(cont container.Container) error {
	var router fiber.Router
	err := cont.Resolve(&router)
	if err != nil {
		return err
	}
	group := router.Group("users")
	group.Get("/", c.GetUsers)
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
