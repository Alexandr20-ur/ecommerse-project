package auth

import "ecommerce/internal/transport"

type Module struct {
	handler *Handler
}

func NewAuthModule(handler *Handler) *Module {
	return &Module{handler: handler}
}

func (m *Module) RegisterRoutes(router transport.Router) {
	router.HandleFunc("POST /auth", m.handler.handlerAuth)
	router.HandleFunc("POST /register", m.handler.handlerRegister)
}
