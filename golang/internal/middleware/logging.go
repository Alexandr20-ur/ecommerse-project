package middleware

import (
	"ecommerce/pkg/log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(logger *log.Logger) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			handler.ServeHTTP(rw, r)

			duration := time.Since(start)

			fields := map[string]any{
				"method":   r.Method,
				"path":     r.URL.Path,
				"status":   rw.statusCode,
				"duration": duration.Milliseconds(),
				"ip":       r.RemoteAddr,
			}

			if rw.statusCode >= 500 {
				logger.ErrorWithFields("Server Error", fields)
			} else if rw.statusCode >= 400 {
				logger.Warn("Client Error")
			} else {
				logger.InfoWithFields("HTTP Request", fields)
			}
		})
	}
}
