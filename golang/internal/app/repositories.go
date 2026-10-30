package app

import (
	"ecommerce/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	User *repository.UserRepository
}

func newRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: repository.NewUserRepository(pool),
	}
}
