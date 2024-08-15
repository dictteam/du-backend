package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/4strodev/du_backend/internal/api"
	"github.com/4strodev/du_backend/internal/api/features/healthcheck"
	"github.com/4strodev/du_backend/internal/api/features/users"
	"github.com/golobby/container/v3"
)

func main() {
	diContainer := container.New()

	diContainer.Singleton(func() *slog.Logger {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})

	// Here dependencies will be attached to the container
	app := api.NewApp(diContainer)

	app.AddProtocolAdapter(&api.HumaProtocolAdapter{})

	err := addControllers(app, healthcheck.NewController(), users.NewUsersController())
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Start())
}

func addControllers(app *api.App, controllers ...api.Controller) error {
	for _, controller := range controllers {
		err := app.AddController(controller)
		if err != nil {
			return err
		}
	}

	return nil
}
