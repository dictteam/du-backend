package api

import (
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/golobby/container/v3"
)

type HumaProtocolAdapter struct {
	api    huma.API
	router *http.ServeMux
	logger *slog.Logger
}

func (h *HumaProtocolAdapter) PrepareDependencies(cont container.Container) error {
	router := http.NewServeMux()
	h.router = router
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	h.api = api
	container.MustResolve(cont, &h.logger)

	cont.Singleton(func() huma.API {
		return api
	})

	return nil
}

func (h *HumaProtocolAdapter) Start() error {
	h.logger.Info("application started on port 8080")
	return http.ListenAndServe("0.0.0.0:8080", h.router)
}
