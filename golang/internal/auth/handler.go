package auth

import (
	"ecommerce/internal/response"
	"ecommerce/pkg/log"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	logger  *log.Logger
	service *Service
}

func NewAuthHandler(logger *log.Logger, authService *Service) *Handler {
	return &Handler{logger, authService}
}

// Авторизация пользователя
func (h *Handler) handlerAuth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var userForm UserForm
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

	var userForm UserForm

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
