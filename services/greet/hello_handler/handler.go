package hello_handler

import (
	"encoding/json"
	"net/http"

	"github.com/eugbyte/monorepo/libs/store/vault"
	"github.com/eugbyte/monorepo/services/webnotify/config"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(rw, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	vaultService := vault.New("https://kv-notify-secrets-stg-ea.vault.azure.net/")
	var fetchVal config.FetchVal = vaultService.GetSecret
	secrets, err := config.FetchAll(fetchVal, "vapid-private-key")
	var privateKey string = secrets[0]

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if privateKey == "" {
		http.Error(rw, "privateKey is empty", http.StatusBadGateway)
		return
	}

	responseBody := map[string]string{
		"stage":   config.Stage().String(),
		"message": "Hello World",
	}

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
