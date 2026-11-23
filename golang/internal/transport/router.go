package transport

import "net/http"

type Router interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type Module interface {
	RegisterRoutes(r Router)
}
