package keyvault

import (
	"errors"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middlewares"
	"github.com/web-notify/api/monorepo/libs/store/vault"
)

type keyVaultMiddleWare struct {
	vaultService vault.VaultServicer
}

func NewKeyVaultMiddleware(vaultService vault.VaultServicer) keyVaultMiddleWare {
	mw := keyVaultMiddleWare{}
	mw.vaultService = vaultService
	return mw
}

func (mw keyVaultMiddleWare) Wrap(handler middlewares.Handler) middlewares.Handler {
	return func(rw http.ResponseWriter, req *http.Request) {
		// pre-process request here
		secretName := "Company-Notification-Key"
		key := req.Header.Get(secretName)
		checkVal, err := mw.vaultService.GetSecret(secretName)
		if err != nil || key != checkVal {
			http.Error(rw, errors.New("Invalid Notification Key").Error(), http.StatusUnauthorized)
			return
		}

		handler(rw, req)

		// post-process reponse here
	}
}
