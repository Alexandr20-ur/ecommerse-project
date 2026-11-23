package auth

import (
	"context"
	"ecommerce/internal/errors"
	"ecommerce/pkg/log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	GenerateToken(userID uuid.UUID) (string, error)
}

type Service struct {
	repository *UserRepository
	logger     *log.Logger
	tokens     TokenGenerator
}

func NewAuthService(repository *UserRepository, logger *log.Logger, tokens TokenGenerator) *Service {
	return &Service{repository: repository, logger: logger, tokens: tokens}
}

func (s Service) Login(ctx context.Context, userForm *UserForm) (string, error) {
	user, err := s.repository.FindByEmail(ctx, userForm.Email)
	if err != nil {
		return "", err //Ошибка базы данных(500)
	}

	if user == nil {
		return "", errors.ErrInvalidCredentials
	}

	//Проверка пароля на актуальность
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userForm.Password))
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	token, err := s.tokens.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s Service) Register(ctx context.Context, userForm *UserForm) error {
	user, err := s.repository.FindByEmail(ctx, userForm.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.ErrUserAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userForm.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user = NewUser(userForm.Email, string(hashPassword))

	err = s.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
