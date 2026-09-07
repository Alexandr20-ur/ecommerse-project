package home

import (
	"ecommerce/internal/handler"
	"ecommerce/pkg/log"
	"fmt"
	"net/http"
)

type HomeHandler struct {
	handler.Handler
}

func NewHomeHandler(logger *log.Logger) (*HomeHandler, error) {
	h, err := handler.NewHandler(logger)
	if err != nil {
		logger.Warn("Failed to create new home handler:" + err.Error())
		return nil, err
	}

	return &HomeHandler{*h}, nil
}
func RegisterRoutes(mux *http.ServeMux, logger *log.Logger) {
	h, err := NewHomeHandler(logger)
	if err != nil {
		logger.Warn("Failed to create new home handler:" + err.Error())
		return
	}
	mux.HandleFunc("GET /", h.handlerMain)
}

func (h *HomeHandler) handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ответ сервера: Home page")
}
