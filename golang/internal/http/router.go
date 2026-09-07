package http

import (
	"ecommerce/internal/home"
	"ecommerce/internal/auth"
	"ecommerce/pkg/log"
	"net/http"
)

func NewRouter(logger *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	home.RegisterRoutes(mux, logger)
	auth.RegisterRoutes(mux, logger)
	return mux
}
