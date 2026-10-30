package database

import (
	"context"
	"ecommerce/internal/config"
	"ecommerce/pkg/log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, config config.DBConfig, logger *log.Logger) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		logger.Error("Failed to parse config:" + err.Error())
		return nil, err
	}

	if cfg.MaxConns > 0 {
		cfg.MaxConns = config.MaxConns
	}
	if cfg.MinConns > 0 {
		cfg.MinConns = config.MinConns
	}
	if cfg.MaxConnIdleTime > 0 {
		cfg.MaxConnIdleTime = config.MaxConnIdleTime
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		logger.Error("Failed to create new pool:" + err.Error())
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Error("Failed to ping pool:" + err.Error())
		pool.Close()
		return nil, err
	}

	return pool, nil
}
