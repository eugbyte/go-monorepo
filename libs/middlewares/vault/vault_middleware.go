package vault

import (
	"errors"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/store/vault"
)

type Handler func(http.Handler) http.Handler

func vaultMiddleware(vaultService vault.VaultServicer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		secretName := "Company-Notification-Key"
		key := req.Header.Get(secretName)
		checkVal, err := vaultService.GetSecret(secretName)
		if err != nil || key != checkVal {
			http.Error(rw, errors.New("Invalid Notification Key").Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(rw, req)
	})
}

func VaultMiddleware(vaultService vault.VaultServicer, next http.Handler) Handler {
	return func(next http.Handler) http.Handler {
		// Dependency injection
		return vaultMiddleware(vaultService, next)
	}
}
