package vault

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middlewares"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type Info struct {
	Username string `json:"username"`
	Company  string `json:"company"`
}

func vaultMiddleware(vaultService vault.VaultServicer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// To be able to read request body multiple times
		body, err := io.ReadAll(req.Body)

		var info Info
		err = json.Unmarshal(body, &info)
		defer req.Body.Close()

		key := req.Header.Get("Notification-Key")
		formats.Trace(info.Company, key)
		checkVal, err := vaultService.GetSecret(info.Company)

		if err != nil {
			formats.Trace(err.Error())
		}
		if err != nil || key != checkVal {
			http.Error(rw, errors.New("Invalid Notification Key").Error(), http.StatusUnauthorized)
			return
		}

		formats.Trace("validation passed")

		// Replace the body with a new reader after reading from the original
		req.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(rw, req)
	})
}

func VaultMiddleware(vaultService vault.VaultServicer) middlewares.HandlerWrapper {
	return func(next http.Handler) http.Handler {
		// Dependency injection
		return vaultMiddleware(vaultService, next)
	}
}
