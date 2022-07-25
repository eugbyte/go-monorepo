package middlewares

import "net/http"

type Handler = func(res http.ResponseWriter, req *http.Request)

type MiddlewareImpl interface {
	Wrap(handler Handler) Handler
}

func Middy(handler Handler, middlewares ...MiddlewareImpl) Handler {
	current := handler
	for _, mw := range middlewares {
		current = mw.Wrap(current)
	}

	return func(rw http.ResponseWriter, request *http.Request) {
		current(rw, request)
	}
}
