package handler

import (
	"ecommerce/pkg/log"
	"net/http"
)

type HandlerInterface interface {
	handlerMain(http.ResponseWriter, *http.Request)
}

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) (*Handler, error) {

	return &Handler{logger: logger}, nil
}

// func NewHandler(logger *log.Logger, tmpl *templates.Store, files ...string) (*Handler, error) {
// 	newTemplate, err := tmpl.NewWithContent(files...)
// 	if err != nil {
// 		logger.Error("Failed to parse template")
// 		return nil, err
// 	}

// 	return &Handler{tmpl: newTemplate, logger: logger}, nil
// }


func (h *Handler) GetLogger() *log.Logger {
	return h.logger
}
