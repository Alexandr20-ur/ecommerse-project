package app

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/config"
	"ecommerce/internal/jwt"
	"ecommerce/internal/product"
	"ecommerce/pkg/log"
)

type Services struct {
	Auth    *auth.Service
	Product *product.Service
}

func newServices(repos *Repositories, cfg *config.Config, logger *log.Logger, jwtManager *jwt.JWTManager) *Services {
	return &Services{
		Auth:    auth.NewAuthService(repos.User, logger, jwtManager),
		Product: product.NewService(repos.Product, logger),
	}
}
