package product

import (
	"context"
	"ecommerce/internal/product/domain"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r Repository) Create(ctx context.Context, product *domain.Product) error {
	query := `INSERT INTO products (id, name, description, price, stock, created_at, updated_at) 
				values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.pool.Exec(
		ctx,
		query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Update(ctx context.Context, product *domain.Product) error {
	return errors.New("not implemented")
}
