package auth

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middleware"
)

type IsAuth func(header http.Header) (bool, error)

// authMiddleware is an example of a middleware layer. It handles the request authorization
// by checking for a key in the url.
func authMiddleware(isAuth IsAuth, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		isValid, err := isAuth(req.Header)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		if isValid {
			next.ServeHTTP(rw, req)
		} else {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func AuthMiddleware(isAuth IsAuth) middleware.HandlerWrapper {
	return func(next http.Handler) http.Handler {
		return authMiddleware(isAuth, next)
	}
}
