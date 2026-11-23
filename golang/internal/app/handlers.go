package app

import (
	"ecommerce/internal/auth"
	myhttp "ecommerce/internal/http"
	"ecommerce/internal/product"
	"ecommerce/pkg/log"
)

func newHandlers(services *Services, logger *log.Logger) *myhttp.Handlers {
	return &myhttp.Handlers{
		Auth:    auth.NewAuthHandler(logger, services.Auth),
		Product: product.NewHandler(logger, services.Product),
	}
}
