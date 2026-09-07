package main

import (
	"context"
	"ecommerce/internal/config"
	myhttp "ecommerce/internal/http"
	"ecommerce/pkg/database"
	mylog "ecommerce/pkg/log"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	logger, err := mylog.New("app.log")
	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig(logger)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := database.NewPool(ctx, cfg.DB, logger)
	defer pool.Close()

	fmt.Println(cfg)

	mux := myhttp.NewRouter(logger)
	handler := corsMiddleware(mux)

	// fs := http.FileServer(http.Dir("./assets"))
	// mux.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// CORS мидлвар
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
