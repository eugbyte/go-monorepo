package hello_handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
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

	httpClient := vault.InsecureClient()
	client := azsecrets.NewClient("https://localhost:8443",
		&vault.FakeCredential{},
		&policy.ClientOptions{Transport: &httpClient})

	secretName := "name"
	secretValue := "Tom"
	params := azsecrets.SetSecretParameters{Value: &secretValue}
	_, err := client.SetSecret(context.TODO(), secretName, params, nil)
	if err != nil {
		log.Fatalf("failed to create a secret: %v", err)
	}

	var ans azsecrets.GetSecretResponse
	ans, err = client.GetSecret(context.TODO(), "name", "", nil)
	formats.Trace(ans.Value)

	responseBody := map[string]interface{}{"message": "Hello World"}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
