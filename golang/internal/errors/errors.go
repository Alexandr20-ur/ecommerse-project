package errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrDatabaseError      = errors.New("database error")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
