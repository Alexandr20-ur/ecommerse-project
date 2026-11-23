package app

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/product"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	User    *auth.UserRepository
	Product *product.Repository
}

func newRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User:    auth.NewUserRepository(pool),
		Product: product.NewRepository(pool),
	}
}
