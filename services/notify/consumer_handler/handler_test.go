package consumer_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mongolib "github.com/web-notify/api/monorepo/libs/db/mongo_lib"
	webpush "github.com/web-notify/api/monorepo/libs/notification/web_push"
	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockMonogoService struct{}

func (ms *MockMonogoService) DB() *mongo.Database {
	return &mongo.Database{}
}
func (ms *MockMonogoService) CreatedShardedCollection(collectionName string, field string, unique bool) {
}
func (ms *MockMonogoService) CreateIndex(collectionName string, field string, unique bool) error {
	return nil
}
func (ms *MockMonogoService) Find(collectionName string, filter primitive.D, items []interface{}) error {
	return nil
}
func (ms *MockMonogoService) FindOne(collectionName string, filter primitive.D, item interface{}) error {
	subscription := models.Subscription{}
	objBytes, err := json.Marshal(subscription)
	if err != nil {
		return err
	}
	err = json.Unmarshal(objBytes, item)
	return err
}
func (ms *MockMonogoService) InsertOne(collectionName string, item interface{}) error {
	return nil
}
func (ms *MockMonogoService) UpdateOne(collectionName string, filter primitive.D, item interface{}, upsert bool) error {
	return nil
}

type MockWebService struct{}

func (ms *MockWebService) SendNotification(message interface{}, endpoint string, auth string, p256dh string, ttl int) error {
	return nil
}

func TestHandler(t *testing.T) {
	info := models.MessageInfo{
		UserID:  "abc@m.com",
		Company: "fakepanda",
		Notification: models.Notification{
			Title: "My title",
			Body:  "My message",
			Icon:  "My icon",
		},
	}

	infoStr := formats.Stringify(info)
	jsonStr := formats.Stringify(infoStr) // for some reason, the azure queue stringyfies the UTF-8 message twice
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
	var mockMongoService mongolib.MonogoServicer = &MockMonogoService{}
	var mockWebService webpush.WebPushServicer = &MockWebService{}

	handler(mockWebService, mockMongoService, writer, request)
	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	fmt.Println(formats.Stringify(data))

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

	isMessageMatch = responseBody.Logs[1] == infoStr
	if !isMessageMatch {
		t.Fatalf("test failed. Expected %v, received %v", message, responseBody.Logs[1])
	}

	t.Logf("test passed")
}
