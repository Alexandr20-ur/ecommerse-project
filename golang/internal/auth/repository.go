package auth

import (
	"context"
	"ecommerce/internal/auth/domain"
	apperrors "ecommerce/internal/errors"
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
		user.PasswordHash,
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

// FindByEmail Получаем пользователя по почте
func (r UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.findOne(ctx, "email", email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID Получаем пользователя по ID
func (r UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := r.findOne(ctx, "id", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) findOne(ctx context.Context, field string, value any) (*domain.User, error) {
	builder := database.Query().
		Select().
		From("users").
		Where(field, "=", value)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var u domain.User

	err = r.pool.QueryRow(ctx, query, args...).Scan(
		&u.ID,
		&u.Email,
		&u.Phone,
		&u.PasswordHash,
		&u.FullName,
		&u.AvatarURL,
		&u.Role,
		&u.IsActive,
		&u.IsEmailVerified,
		&u.LastLoginAt,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrUserNotFound
		}
		return nil, err
	}

	return &u, nil
}
