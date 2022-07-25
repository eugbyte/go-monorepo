package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type RequestBody struct {
	Num int `json:"num"`
}

func MockHandler(rw http.ResponseWriter, request *http.Request) {
	myVal := request.Header.Get("my_value")

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(rw).Encode(myVal)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Mock middleware that appends a value to the header
type MockMiddleWare struct {
}

func (mw MockMiddleWare) Wrap(handler Handler) Handler {
	return func(rw http.ResponseWriter, req *http.Request) {
		// let's extract the number from the request body, and multiply it by two
		formats.Trace("LogMiddleware", "pre-processing request...")
		req.Header.Set("my_value", "my_header")

		handler(rw, req)
	}
}

func Test_Middy(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	mockMiddleware := MockMiddleWare{}
	handler := Middy(MockHandler, mockMiddleware)

	writer := httptest.NewRecorder()
	handler(writer, request)
	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	var headerValue string
	err = json.Unmarshal(data, &headerValue)
	if err != nil {
		t.Fatal("Error unmarshalling response body to map[string]string")
	}

	t.Log(headerValue)
	if headerValue != "my_header" {
		t.Fatalf("test failed. Expected %v, received %v", "my_header", headerValue)
	}
}
