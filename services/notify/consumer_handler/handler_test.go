package consumer_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
)

func TestHandler(t *testing.T) {
	jsonStr := "\"{\\\"endpoint\\\":\\\"http://localhost:3000\\\",\\\"message\\\":\\\"your order has been received\\\",\\\"expirationTime\\\":\\\"1657034161512\\\",\\\"keys\\\":{\\\"auth\\\":\\\"123\\\",\\\"p256dh\\\":\\\"abc\\\"}}\""
	reqData := map[string]string{
		"req": jsonStr,
	}
	requestBody := map[string]interface{}{
		"Data":     reqData,
		"Metadata": map[string]string{},
	}
	objBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("Cannot marshal", err.Error())
	}

	request := httptest.NewRequest(http.MethodPost, "/hello", bytes.NewBuffer(objBytes))

	writer := httptest.NewRecorder()
	Handler(writer, request)
	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	var responseBody qmodels.ResponseBody
	err = json.Unmarshal(data, &responseBody)
	if err != nil {
		t.Fatal("Error unmarshalling response body to qmodels.ResponseBody")
	}

	var message string
	err = json.Unmarshal([]byte(jsonStr), &message)
	if err != nil {
		t.Fatal("Error unmarshalling response body to qmodels.ResponseBody")
	}

	isMessageMatch := responseBody.Logs[0] == "Message successfully dequeued"
	if !isMessageMatch {
		t.Fatalf("test failed. Expected %v, received %v", "Message successfully dequeued", responseBody.Logs[0])
	}

	isMessageMatch = responseBody.Logs[1] == fmt.Sprintf("message: '%s'", message)
	if !isMessageMatch {
		t.Fatalf("test failed. Expected %v, received %v", message, responseBody.Logs[1])
	}

	t.Logf("test passed")

}
