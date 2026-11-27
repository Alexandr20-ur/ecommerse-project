package auth

import (
	"ecommerce/internal/auth/dto"
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
	var userForm dto.UserForm
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

	var userForm dto.UserForm

	err := json.NewDecoder(r.Body).Decode(&userForm)

	if err != nil {
		h.logger.Error("Failder to parse body:" + err.Error())
		response.WriteJson(w, http.StatusBadRequest, "Failed to parse body:"+err.Error())
		return
	}

	err = h.service.Register(r.Context(), &userForm)
	if err != nil {
		h.logger.Error("Failder to register user:" + err.Error())
		response.WriteJson(w, http.StatusUnauthorized, "Failed to register user:"+err.Error())
		return
	}

	fmt.Fprintf(w, "OK")
}

// Проверка полученного токена пользователя
func (h *Handler) handlerCheckJWT(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		h.logger.Error("Failder to parse token")
		response.WriteJson(w, http.StatusUnauthorized, "Failder to parse token")
		return
	}

	user, err := h.service.CheckToken(r.Context(), jwtToken)
	if err != nil {
		h.logger.Error("Failder to check token:" + err.Error())
		response.WriteJson(w, http.StatusUnauthorized, "Failder to check token")
		return
	}

	response.WriteJson(w, http.StatusOK, dto.NewUserResponse(user))
}
