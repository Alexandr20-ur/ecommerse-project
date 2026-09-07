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
		logger.Warn("Failed to parse config:" + err.Error())
		return nil, err
	}

	cfg.MaxConns = int32(config.MaxOpenConns)

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		logger.Warn("Failed to create new pool:" + err.Error())
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		logger.Warn("Failed to ping pool:" + err.Error())
		return nil, err
	}

	return pool, nil
}
