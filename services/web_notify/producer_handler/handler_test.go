package producerhandler

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	qlib "github.com/eugbyte/monorepo/libs/queue"
	"github.com/eugbyte/monorepo/services/webnotify/models"
)

type MockQueueService struct{}

func (qService *MockQueueService) Init(cxt context.Context, queueName string, rootConnection string, accountName string, accountKey string) {
}
func (qService *MockQueueService) QueueExist() bool {
	return true
}
func (qService *MockQueueService) CreateQueue(metaData azqueue.Metadata) error {
	return nil
}
func (qService *MockQueueService) Enqueue(messageText string, visibilityTimeout time.Duration, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error) {
	resp := azqueue.EnqueueMessageResponse{}
	return &resp, nil
}

func TestHandler(t *testing.T) {
	var mockQueueService qlib.QueueServicer = &MockQueueService{}

	mockReq := models.MessageInfo{
		Company: "fakepanda",
		UserID:  "abc@m.com",
	}

	objBytes, err := json.Marshal(mockReq)
	if err != nil {
		t.Fatal("Cannot marshal", err.Error())
	}
	request := httptest.NewRequest(http.MethodPost, "/api/notifications", bytes.NewBuffer(objBytes))

	writer := httptest.NewRecorder()
	handler(mockQueueService, writer, request)
	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	var responseBody map[string]string
	err = json.Unmarshal(data, &responseBody)
	if err != nil {
		t.Fatal("Error unmarshalling response body to map[string]string")
	}

	message := responseBody["message"]
	if message != "successfully enqueued" {
		t.Fatalf("expected '%v', but received '%v'", "successfully enqueued", message)
	}
}
