package product

import (
	"ecommerce/internal/product/domain"
	"ecommerce/internal/response"
	"ecommerce/pkg/log"
	"encoding/json"
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

	var productForm domain.ProductForm
	var resp response.ProductResponse
	err := json.NewDecoder(r.Body).Decode(&productForm)

	if err != nil {
		resp.Err = "Failder to parse body:" + err.Error()
		h.logger.Error(resp.Err)
		response.WriteJson(w, http.StatusBadRequest, resp)
		return
	}

	h.service.Create(r.Context(), &productForm)

	response.WriteJson(w, http.StatusOK, resp)
}

func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ShowHandler(w http.ResponseWriter, r *http.Request) {}
