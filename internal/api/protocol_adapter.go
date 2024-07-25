package api

import "github.com/golobby/container/v3"

type ProtocolAdapter interface {
	PrepareDependency(cont container.Container) error
	Start() error
}
