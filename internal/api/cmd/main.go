package main

import (
	"log"

	"github.com/4strodev/du_backend/internal/api"
	"github.com/4strodev/du_backend/internal/api/features/healthcheck"
	"github.com/4strodev/du_backend/internal/api/features/users"
	"github.com/golobby/container/v3"
)

func main() {
	diContainer := container.New()

	// Here dependencies will be attached to the container
	app := api.NewApp(diContainer)

	app.AddProtocolAdapter(&api.FiberProtocolResolver{})

	err := app.AddController(healthcheck.NewController())
	if err != nil {
		panic(err)
	}
	err = app.AddController(users.NewUsersController())
	if err != nil {
		panic(err)
	}

	log.Fatal(app.Start())
}
