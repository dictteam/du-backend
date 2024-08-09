package api

import (
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/golobby/container/v3"
)

type HumaProtocolResolver struct {
	api    huma.API
	router *http.ServeMux
}

func (h *HumaProtocolResolver) PrepareDependencies(cont container.Container) error {
	router := http.NewServeMux()
	h.router = router
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	h.api = api

	cont.Singleton(func() huma.API {
		return api
	})

	return nil
}

func (h *HumaProtocolResolver) Start() error {
	log.Println("starting http server on port 8080")
	return http.ListenAndServe("0.0.0.0:8080", h.router)
}
