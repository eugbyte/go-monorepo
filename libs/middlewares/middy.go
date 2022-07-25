package middlewares

import "net/http"

type Handler = func(http.Handler) http.Handler

// Middleware pattern from https://golangcode.com/middleware-on-handlers .
// Middy makes adding more than one layer of middleware easy by specifying them as a list
func Middy(handler http.Handler, middlewares ...Handler) http.Handler {
	current := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		var mw Handler = middlewares[i]
		current = mw(current)
	}
	return current
}
