package users

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	fiber "github.com/gofiber/fiber/v2"
	container "github.com/golobby/container/v3"
)

type UserController struct {
}

type UserResponse struct {
	Body struct {
		Message string `json:"message" doc:"message received from api" example:"Hello from users" requied:"false"`
	}
}

// Init implements api.Controller.
func (c *UserController) Init(cont container.Container) error {
	var api huma.API
	err := cont.Resolve(&api)
	if err != nil {
		return err
	}

	huma.Get(api, "/users", func(ctx context.Context, body *struct{}) (*UserResponse, error) {
		response := &UserResponse{}
		response.Body.Message = "Hello from users"
		return response, nil
	})

	return nil
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"msg": "Method yet not implemented",
	})
}

// ShutDown implements api.Controller.
func (*UserController) ShutDown() error {
	return nil
}

func NewUsersController() *UserController {
	return &UserController{}
}
