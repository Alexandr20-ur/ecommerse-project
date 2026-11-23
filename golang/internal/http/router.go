package http

import (
	"ecommerce/internal/transport"
	"net/http"
)

func NewRouter(modules ...transport.Module) *http.ServeMux {
	mux := http.NewServeMux()

	apiMux := http.NewServeMux()

	for _, module := range modules {
		module.RegisterRoutes(apiMux)
	}

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))
	return mux
}
