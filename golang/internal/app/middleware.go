package app

import (
	"ecommerce/internal/jwt"
	"ecommerce/internal/middleware"
)

func publicMiddlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func protectedMiddlewares(jwtManager *jwt.JWTManager) []middleware.Middleware {
	return []middleware.Middleware{
		middleware.JWT(jwtManager),
	}
}

func globalMiddlewares() []middleware.Middleware {
	return []middleware.Middleware{
		middleware.CORS(),
	}
}
