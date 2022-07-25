package logger

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

// AuthMiddleware is an example of a middleware layer. It handles the request authorization
// by checking for a key in the url.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		formats.Trace("preprocessing request...")

		next.ServeHTTP(rw, req)
	})
}
