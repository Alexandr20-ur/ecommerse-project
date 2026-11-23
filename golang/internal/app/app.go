package app

import (
	"context"
	"ecommerce/internal/config"
	myhttp "ecommerce/internal/http"
	"ecommerce/internal/jwt"
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

	jwtManager := jwt.NewJWTManager(cfg.JWT.Secret, cfg.JWT.TTL)

	repos := newRepositories(pool)
	services := newServices(repos, cfg, logger, jwtManager)
	handlers := newHandlers(services, logger)

	modules := NewModules(handlers)
	mux := myhttp.NewRouter(modules.All()...)
	return serve(mux)
}

func serve(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         ":8000",
		Handler:      middleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
