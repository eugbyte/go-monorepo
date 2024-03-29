package samplepushhandler

import (
	"net/http"

	"github.com/eugbyte/monorepo/libs/middleware"
)

// Dependency injection

var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	client := http.Client{}
	handler(&client, rw, req)
})

// Wrap middlewares

var HTTPHandler http.Handler = middleware.Middy(httpHandler)
