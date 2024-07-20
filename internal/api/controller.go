package api

import "github.com/golobby/container/v3"

type Controller interface {
	// Init Initialize controller dependencies
	Init(cont container.Container)
	// ShutDown gracefully shutdown controller and current transactions
	ShutDown()
}
