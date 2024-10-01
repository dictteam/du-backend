package healthcheck

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	container "github.com/golobby/container/v3"
)

type HealthCheckController struct {
}

type HealtcheckOutput struct {
	ContentType string `header:"Content-Type"`
	Body []byte
}

// Init implements api.Controller.
func (*HealthCheckController) Init(cont container.Container) error {
	var api huma.API
	err := cont.Resolve(&api)
	if err != nil {
		return err
	}

	huma.Get(api, "/healthcheck", func(ctx context.Context, input *struct {}) (*HealtcheckOutput, error) {
		resp := &HealtcheckOutput{}
		resp.ContentType = "text/plain"
		resp.Body = []byte("ok")
		return resp, nil
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
