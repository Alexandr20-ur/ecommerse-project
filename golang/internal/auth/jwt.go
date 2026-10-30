package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager struct {
	secretKey []byte
	ttl       time.Duration
}

type JwtClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken Создание JWT токена
func (m *JWTManager) GenerateToken(userID uuid.UUID) (string, error) {
	claims := &JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

// NewJWTManager Construct
func NewJWTManager(secretKey string, ttl time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secretKey),
		ttl:       ttl,
	}
}
