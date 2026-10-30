package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	Email           string     `json:"email" db:"email"`
	Phone           *string    `json:"phone,omitempty" db:"phone"` // указатель, т.к. может быть NULL
	PasswordHash    string     `json:"-" db:"password_hash"`
	FullName        string     `json:"full_name" db:"full_name"`
	AvatarURL       *string    `json:"avatar_url,omitempty" db:"avatar_url"`
	Role            string     `json:"role" db:"role"`
	IsActive        bool       `json:"is_active" db:"is_active"`
	IsEmailVerified bool       `json:"is_email_verified" db:"is_email_verified"`
	LastLoginAt     *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
}

type UserForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUser(login string, hashPassword string) *User {
	now := time.Now()
	return &User{
		ID:              uuid.New(),
		Email:           login,
		PasswordHash:    hashPassword,
		Role:            "user",
		IsActive:        true,
		IsEmailVerified: false,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
