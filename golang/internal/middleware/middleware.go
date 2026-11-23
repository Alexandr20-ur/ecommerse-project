package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Chain применяет middleware так, что первый в списке становится - внешний
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
