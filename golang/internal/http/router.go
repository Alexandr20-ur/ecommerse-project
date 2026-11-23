package http

import (
	"ecommerce/internal/middleware"
	"ecommerce/internal/transport"
	"net/http"
)

func NewRouter(
	modules []transport.Module,
	pub, prot []middleware.Middleware,
) *http.ServeMux {
	mux := http.NewServeMux()

	apiMux := http.NewServeMux()
	publicMux := http.NewServeMux()

	for _, module := range modules {
		module.RegisterRoutes(apiMux)
		module.RegisterPublicRoutes(publicMux)
	}

	publicHandler := middleware.Chain(publicMux, pub...)
	protectedHandler := middleware.Chain(apiMux, prot...)

	mux.Handle("/api/", http.StripPrefix("/api", protectedHandler))
	mux.Handle("/", publicHandler)

	return mux
}
