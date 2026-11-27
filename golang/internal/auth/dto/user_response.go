package dto

import (
	"ecommerce/internal/auth/domain"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone,omitempty"`
	FullName  string    `json:"full_name"`
	AvatarURL *string   `json:"avatar_url,omitempty"`
	Role      string    `json:"role"`
}

// Маппинг из домена в DTO
func NewUserResponse(u *domain.User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Phone:     u.Phone,
		FullName:  u.FullName,
		AvatarURL: u.AvatarURL,
		Role:      u.Role,
	}
}
