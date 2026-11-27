package auth

import "ecommerce/internal/transport"

type Module struct {
	handler *Handler
}

func NewAuthModule(handler *Handler) *Module {
	return &Module{handler: handler}
}

func (m *Module) RegisterPublicRoutes(router transport.Router) {
	router.HandleFunc("POST /auth", m.handler.handlerAuth)
	router.HandleFunc("POST /register", m.handler.handlerRegister)
	router.HandleFunc("GET /me", m.handler.handlerCheckJWT) //Тут выполняем проверку на валидность JWT токена. Что пользователь,не мог зайти в систему с протухшим токеном
}

func (m *Module) RegisterRoutes(router transport.Router) {

}
