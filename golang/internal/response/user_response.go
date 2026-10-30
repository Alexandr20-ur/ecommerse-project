package response

import (
	"ecommerce/internal/domain"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	Phone           *string   `json:"phone,omitempty"`
	FullName        string    `json:"full_name"`
	AvatarURL       *string   `json:"avatar_url,omitempty"`
	Role            string    `json:"role"`
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewUserResponse(u *domain.User) *UserResponse {
	return &UserResponse{
		ID:              u.ID,
		Email:           u.Email,
		Phone:           u.Phone,
		FullName:        u.FullName,
		AvatarURL:       u.AvatarURL,
		Role:            u.Role,
		IsEmailVerified: u.IsEmailVerified,
		CreatedAt:       u.CreatedAt,
	}
}
