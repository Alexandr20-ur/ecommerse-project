package jwt

import (
	"ecommerce/internal/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager struct {
	secretKey []byte
	config    *config.JWTConfig
}

type JwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken Создание JWT токена
func (m *JWTManager) GenerateToken(userID uuid.UUID) (string, error) {
	now := time.Now()

	claims := &JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(m.config.AccessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (m *JWTManager) ParseToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}

	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return m.secretKey, nil
		},
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithLeeway(m.config.Leeway),
	)

	if err != nil {
		return nil, err
	}

	return claims, nil
}

// NewJWTManager Construct
func NewJWTManager(cfg *config.JWTConfig) (*JWTManager, error) {
	lenSecret := len(cfg.Secret)
	if lenSecret < 32 {
		return nil, fmt.Errorf("jwt: secret key must be at least 32 bytes, got %d", lenSecret)
	}

	if cfg.AccessTTL <= 0 {
		return nil, fmt.Errorf("jwt: ttl must be positive, got %v", cfg.AccessTTL)
	}

	return &JWTManager{
		secretKey: []byte(cfg.Secret),
		config:    cfg,
	}, nil
}
