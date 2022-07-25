package vault

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/web-notify/api/monorepo/libs/middlewares"
)

func Test_VaultMW(t *testing.T) {
	mockVaultService := mockVaultService{}
	vaultVM := VaultMiddleware(mockVaultService)
	var wrappedHandler http.Handler = middlewares.Middy(http.HandlerFunc(handler), vaultVM)

	request := httptest.NewRequest(http.MethodPost, "/hello", nil)

	writer := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(writer, request)

	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	var messageMap map[string]string
	err = json.Unmarshal(data, &messageMap)
	if err != nil {
		t.Fatal("Error unmarshalling response body to map[string]string")
	}
	message := messageMap["message"]
	if message != "Hello World" {
		t.Fatalf("test failed. Expected %v, received %v", "Hello World", message)
	}
}

func handler(response http.ResponseWriter, request *http.Request) {
	responseBody := map[string]interface{}{"message": "Hello World"}
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

type mockVaultService struct {
}

func (vs mockVaultService) GetSecret(secretName string) (string, error) {
	return "abc", nil
}
func (vs mockVaultService) SetSecret(secretName string, secretValue string) error {
	return nil
}
