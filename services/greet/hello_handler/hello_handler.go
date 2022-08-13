package hello_handler

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middleware"
)

// Dependency injection, if any
// e.g. handler(NewService(), rw, req)
var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	handler(rw, req)
})

// Wrap middlewares
// middleware applied here, in contrast to the mux server, will be for this specific controller only
var HTTPHandler http.Handler = middleware.Middy(httpHandler)
