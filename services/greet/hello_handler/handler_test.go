package hello_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	request := httptest.NewRequest(http.MethodPost, "/hello", nil)

	writer := httptest.NewRecorder()
	Handler(writer, request)
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
