package product

import (
	"context"
	"ecommerce/internal/product/domain"
	"ecommerce/pkg/log"
)

type Service struct {
	repository *Repository
	logger     *log.Logger
}

func NewService(repository *Repository, logger *log.Logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) Create(ctx context.Context, productForm *domain.ProductForm) error {
	product := domain.NewProduct(productForm)
	err := s.repository.Create(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
