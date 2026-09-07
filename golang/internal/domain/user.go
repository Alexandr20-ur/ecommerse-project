package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"` //Не выводим в json
	FullName     string
}


type UserForm struct {
	Login string `json:"login"`
	Password string `json:"password"`
}