package product

import (
	"ecommerce/internal/transport"
)

type Module struct {
	handler *Handler
}

func NewProductModule(handler *Handler) *Module {
	return &Module{
		handler: handler,
	}
}

func (m *Module) RegisterRoutes(router transport.Router) {
	router.HandleFunc("POST /products", m.handler.CreateHandler)
	router.HandleFunc("PUT /products/{id}", m.handler.UpdateHandler)
	router.HandleFunc("DELETE /products/{id}", m.handler.DeleteHandler)
	router.HandleFunc("GET /products", m.handler.GetAllHandler)
	router.HandleFunc("GET /products/{id}", m.handler.ShowHandler)
}
