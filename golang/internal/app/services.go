package app

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/config"
	"ecommerce/internal/service"
	"ecommerce/pkg/log"
)

type Services struct {
	Auth *service.AuthService
}

func newServices(repos *Repositories, cfg *config.Config, logger *log.Logger, jwtManager *auth.JWTManager) *Services {
	return &Services{
		Auth: service.NewAuthService(repos.User, logger, jwtManager),
	}
}
