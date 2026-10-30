package service

import (
	"context"
	"ecommerce/internal/domain"
	"ecommerce/internal/repository"
	"ecommerce/pkg/log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerator interface {
	GenerateToken(userID uuid.UUID) (string, error)
}

type AuthService struct {
	repository *repository.UserRepository
	logger     *log.Logger
	tokens     TokenGenerator
}

func NewAuthService(repository *repository.UserRepository, logger *log.Logger, tokens TokenGenerator) *AuthService {
	return &AuthService{repository: repository, logger: logger, tokens: tokens}
}

func (s AuthService) Login(ctx context.Context, userForm *domain.UserForm) (string, error) {
	user, err := s.repository.FindByEmail(ctx, userForm.Login)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", nil
	}

	token, err := s.tokens.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s AuthService) Register(ctx context.Context, userForm *domain.UserForm) error {
	user, err := s.repository.FindByEmail(ctx, userForm.Login)
	if err != nil {
		return err
	}

	if user != nil {
		return domain.ErrUserAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userForm.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user = domain.NewUser(userForm.Login, string(hashPassword))

	err = s.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
