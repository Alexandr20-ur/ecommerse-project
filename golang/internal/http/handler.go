package http

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/product"
)

type Handlers struct {
	Auth    *auth.Handler
	Product *product.Handler
}
