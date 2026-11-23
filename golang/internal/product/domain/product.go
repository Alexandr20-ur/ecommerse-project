package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"` // Указатель, т.к. в БД может быть NULL
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProduct(form *ProductForm) *Product {
	now := time.Now()
	return &Product{
		ID:          uuid.New(),
		Name:        form.Name,
		Description: &form.Description,
		Price:       form.Price,
		Stock:       form.Stock,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
