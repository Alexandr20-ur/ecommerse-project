package app

import (
	"ecommerce/internal/jwt"
	"ecommerce/internal/middleware"
	"ecommerce/pkg/log"
)

func publicMiddlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func protectedMiddlewares(jwtManager *jwt.JWTManager) []middleware.Middleware {
	return []middleware.Middleware{
		middleware.JWT(jwtManager),
	}
}

func globalMiddlewares(logger *log.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		middleware.CORS(),
		middleware.Logger(logger),
	}
}
