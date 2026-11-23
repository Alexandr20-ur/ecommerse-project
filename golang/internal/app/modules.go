package app

import (
	"ecommerce/internal/auth"
	myhttp "ecommerce/internal/http"
	"ecommerce/internal/product"
	"ecommerce/internal/transport"
)

type Modules struct {
	Auth    *auth.Module
	Product *product.Module
}

func newModules(h *myhttp.Handlers) *Modules {
	return &Modules{
		Auth:    auth.NewAuthModule(h.Auth),
		Product: product.NewProductModule(h.Product),
	}
}

func (m *Modules) All() []transport.Module {
	return []transport.Module{
		m.Auth,
		m.Product,
	}
}
