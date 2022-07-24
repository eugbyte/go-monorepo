package hello_handler

import (
	"encoding/json"
	"log"
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

	vs := vault.NewVaultService("https://localhost:8443")
	secretName := "name"
	secretValue := "Tom"
	err := vs.SetSecret(secretName, secretValue)
	if err != nil {
		formats.Trace(err)
		log.Fatalf("failed to create a secret: %v", err)
	}
	ans, err := vs.GetSecret(secretName)
	formats.Trace(ans)

	responseBody := map[string]interface{}{"message": "Hello World"}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
