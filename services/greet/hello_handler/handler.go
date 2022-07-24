package hello_handler

import (
	"encoding/json"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type RequestBody struct {
	Message string `json:"message"`
}

func Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	secretName := "db_Name"
	secretValue := "123"

	vs := vault.NewVaultService("https://localhost:8443")
	err := vs.SetSecret(secretName, secretValue)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	formats.Trace(secretName, secretValue, vs)

	responseBody := map[string]interface{}{"message": "Hello World"}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
