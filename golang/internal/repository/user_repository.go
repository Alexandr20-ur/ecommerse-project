package repository

import (
	"context"
	"ecommerce/internal/domain"
	"ecommerce/pkg/database"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool}
}

func (r UserRepository) Create(ctx context.Context, user *domain.User) error {

	query := `
        INSERT INTO users (
            id, email, phone, password_hash, full_name, 
            avatar_url, role, is_active, is_email_verified,
            created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
        )
    `
	_, err := r.pool.Exec(
		ctx,
		query,
		user.ID,
		user.Email,
		user.Phone,
		user.FullName,
		user.AvatarURL,
		user.Role,
		user.IsActive,
		user.IsEmailVerified,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {

	builder := database.Query().
		Select().
		From("users").
		Where("email", "=", email)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var u domain.User
	err = r.pool.QueryRow(ctx, query, args...).Scan(&u)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // «не найден» — это не ошибка на уровне репозитория
		}
		return nil, err
	}

	return &u, nil
}
