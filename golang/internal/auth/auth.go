package auth

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/response"
	"ecommerce/internal/service"
	"ecommerce/pkg/log"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	logger  *log.Logger
	service *service.AuthService
}

func NewAuthHandler(logger *log.Logger, authService *service.AuthService) *Handler {
	return &Handler{logger, authService}
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /auth", h.handlerAuth)
	mux.HandleFunc("POST /register", h.handlerRegister)
}

// Авторизация пользователя
func (h *Handler) handlerAuth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userForm domain.UserForm
	resp := response.LoginResponse{}

	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		resp.Err = "Failder to parse body:" + err.Error()
		h.logger.Error(resp.Err)
		response.WriteJson(w, http.StatusBadRequest, resp)
		return
	}

	token, err := h.service.Login(r.Context(), &userForm)

	if err != nil {
		resp.Err = "Failder to login:" + err.Error()
		h.logger.Error(resp.Err)
		response.WriteJson(w, http.StatusUnauthorized, resp)
		return
	}

	resp.Token = token
	response.WriteJson(w, http.StatusOK, resp)
	return
}

// Регистрация пользователя
func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userForm domain.UserForm

	err := json.NewDecoder(r.Body).Decode(&userForm)

	if err != nil {
		h.logger.Error("Failder to parse body:" + err.Error())
		http.Error(w, "Failder to parse body:"+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Register(r.Context(), &userForm)
	if err != nil {
		h.logger.Error("Failder to register user:" + err.Error())
		http.Error(w, "Failder to register user:"+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "OK")
}
