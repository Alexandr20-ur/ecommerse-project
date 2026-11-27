package auth

import (
	"context"
	"ecommerce/internal/auth/domain"
	"ecommerce/internal/auth/dto"
	"ecommerce/internal/errors"
	apperrors "ecommerce/internal/errors"
	"ecommerce/internal/jwt"
	"ecommerce/pkg/log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *UserRepository
	logger     *log.Logger
	jwtManager *jwt.JWTManager
}

func NewAuthService(repository *UserRepository, logger *log.Logger, jwtManager *jwt.JWTManager) *Service {
	return &Service{repository: repository, logger: logger, jwtManager: jwtManager}
}

func (s Service) Login(ctx context.Context, userForm *dto.UserForm) (string, error) {
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

	token, err := s.jwtManager.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s Service) Register(ctx context.Context, userForm *dto.UserForm) error {
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

	user = domain.NewUser(userForm.Email, string(hashPassword))

	err = s.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) CheckToken(ctx context.Context, jwtToken string) (*domain.User, error) {
	parts := strings.SplitN(jwtToken, " ", 2)

	// Проверяем, что формат верный: "Bearer <token>"
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, apperrors.ErrInvalidCredentials
	}

	tokenString := parts[1]

	claims, err := s.jwtManager.ParseToken(tokenString)
	if err != nil {
		s.logger.Error("Failder to parse token:" + err.Error())
		return nil, apperrors.ErrInvalidCredentials
	}

	user, err := s.repository.FindByID(ctx, claims.UserID.String())
	if err != nil {
		s.logger.Error("Failder to find user:" + err.Error())
		return nil, apperrors.ErrInvalidCredentials
	}

	return user, nil
}
