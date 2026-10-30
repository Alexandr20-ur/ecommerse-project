package http

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/home"
	"net/http"
)

type Handlers struct {
	Home *home.HomeHandler
	Auth *auth.Handler
}

func NewRouter(h Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	home.RegisterRoutes(mux, h.Home)
	auth.RegisterRoutes(mux, h.Auth)
	return mux
}
