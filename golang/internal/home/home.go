package home

import (
	"ecommerce/pkg/log"
	"fmt"
	"net/http"
)

type HomeHandler struct {
	logger *log.Logger
}

func NewHomeHandler(logger *log.Logger) *HomeHandler {
	return &HomeHandler{logger: logger}
}

func RegisterRoutes(mux *http.ServeMux, h *HomeHandler) {
	mux.HandleFunc("GET /", h.handlerMain)
}

func (h *HomeHandler) handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ответ сервера: Home page")
}
