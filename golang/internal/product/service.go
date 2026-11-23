package product

import "ecommerce/pkg/log"

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
