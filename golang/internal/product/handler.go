package product

import (
	"ecommerce/pkg/log"
	"fmt"
	"net/http"
)

type Handler struct {
	logger  *log.Logger
	service *Service
}

func NewHandler(logger *log.Logger, service *Service) *Handler {
	return &Handler{logger, service}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data := r.Body
	fmt.Printf("%s", data)
}

func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ShowHandler(w http.ResponseWriter, r *http.Request) {}
