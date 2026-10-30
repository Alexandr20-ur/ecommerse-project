package app

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/home"
	myhttp "ecommerce/internal/http"
	"ecommerce/pkg/log"
)

func newHandlers(services *Services, logger *log.Logger) myhttp.Handlers {
	return myhttp.Handlers{
		Home: home.NewHomeHandler(logger),
		Auth: auth.NewAuthHandler(logger, services.Auth),
	}
}
