package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eugbyte/monorepo/libs/middleware"
)

func Test_VaultMW(t *testing.T) {

	request := httptest.NewRequest(http.MethodPost, "/hello", nil)

	writer := httptest.NewRecorder()

	var isAuth IsAuth = func(header http.Header) (bool, error) {
		return false, errors.New("Somethin went wrong")
	}

	var httpHandler = middleware.Middy(http.HandlerFunc(handler), AuthMiddleware(isAuth))

	httpHandler.ServeHTTP(writer, request)

	result := writer.Result()
	defer result.Body.Close()

	_, err := ioutil.ReadAll(result.Body)
	if err == nil {
		t.Fatal("Expected error, but non received")
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
