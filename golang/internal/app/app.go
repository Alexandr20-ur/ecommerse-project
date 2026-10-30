package app

import (
	"context"
	"ecommerce/internal/auth"
	"ecommerce/internal/config"
	myhttp "ecommerce/internal/http"
	"ecommerce/pkg/database"
	mylog "ecommerce/pkg/log"
	"net/http"
	"time"
)

func Run() error {
	logger, err := mylog.New("app.log")
	if err != nil {
		return err
	}

	cfg, err := config.NewConfig(logger)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := database.NewPool(ctx, cfg.DB, logger)
	if err != nil {
		return err
	}

	defer pool.Close()

	jwtManager := auth.NewJWTManager(cfg.JWT.Secret, cfg.JWT.TTL)

	repos := newRepositories(pool)
	services := newServices(repos, cfg, logger, jwtManager)
	handlers := newHandlers(services, logger)

	mux := myhttp.NewRouter(handlers)
	return serve(mux)
}

func serve(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      corsMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем все origins для разработки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Обрабатываем preflight запросы (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
