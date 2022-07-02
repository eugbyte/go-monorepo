package middleware

import "net/http"

type Middleware struct {
	handler http.Handler
}

func (mw *Middleware) SetHandler(handler http.Handler) *Middleware {
	mw.handler = handler
	return mw
}
