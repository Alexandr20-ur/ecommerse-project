package auth

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/handler"
	"ecommerce/pkg/log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AuthHandler struct {
	handler.Handler
}

func newAuthHandler( logger *log.Logger) (*AuthHandler, error) {
	h, err := handler.NewHandler(logger)
	if err != nil {
		logger.Warn("Failed to create new home handler:" + err.Error())
		return nil, err
	}

	return &AuthHandler{*h}, nil
}

func RegisterRoutes(mux *http.ServeMux, logger *log.Logger) {
	h, err := newAuthHandler(logger)
	if err != nil {
		logger.Warn("Failed to create new home handler:" + err.Error())
		return
	}
	mux.HandleFunc("POST /auth", h.handlerAuth)
	mux.HandleFunc("POST /register", h.handlerRegister)
}

//Авторизация пользователя
func (h *AuthHandler) handlerAuth(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.GetLogger().Warn("Failder to read body:" + err.Error())
		http.Error(w, "Failed to read body:" + err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	fmt.Println("Raw body:", string(body))
}

//Регистрация пользователя
func (h *AuthHandler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.GetLogger().Warn(err.Error())
		http.Error(w, "Failder to read body:" + err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var loginData domain.UserForm
	if err := json.Unmarshal(body, &loginData); err != nil {
		h.GetLogger().Warn(err.Error())
		http.Error(w, "Failder to parse json: " + err.Error(), http.StatusBadGateway)
		return
	}

	
}
