package sample_push_handler

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middleware"
)

// Dependency injection
var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	client := http.Client{}
	handler(&client, rw, req)
})

// Wrap middlewares
var HTTPHandler http.Handler = middleware.Middy(httpHandler)
