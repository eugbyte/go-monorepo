package subscriberhandler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/services/webnotify/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockMonogoService struct{}

func (ms *MockMonogoService) Init(dbName string, connectionString string) {}
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
	return nil
}
func (ms *MockMonogoService) InsertOne(collectionName string, item interface{}) error {
	return nil
}
func (ms *MockMonogoService) UpdateOne(collectionName string, filter primitive.D, item interface{}, upsert bool) error {
	return nil
}

func TestHandler(t *testing.T) {
	mockMongoService := MockMonogoService{}

	mockReq := models.Subscription{
		Company:  "fakepanda",
		UserID:   "abc@m.com",
		Endpoint: "https://fcm.googleapis.com/fcm/send/ePpws-p5uBI:APA91bFm_zyeVqFGxiw5kWJJR9KLy9eFRRUKbc8_yebfBbsNBcX3iJmAUgl22uu_zpe2Hv0RSHpEThblr_Bz5AbbHQ7WXuUO2LkxmqJjTc6N1GURuZEOjtt2Y6_pr0org506K1ZMk6lK",
		Keys: models.Keys{
			P256dh: "BKOynOSa_eUI1ZGhmWxsaA34lbfqtxGTXiZTFa24SpDjOQBHwwCfxLBuWRdO_92E5A3ia8VA3Q5774ECPK6-Khg",
			Auth:   "wquys90eGkBzfmTSyMT-PQ",
		},
	}

	objBytes, err := json.Marshal(mockReq)
	if err != nil {
		t.Fatal("Cannot marshal", err.Error())
	}
	request := httptest.NewRequest(http.MethodPost, "/api/subscriptions", bytes.NewBuffer(objBytes))

	writer := httptest.NewRecorder()
	handler(&mockMongoService, writer, request)
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

	formats.Trace(responseBody)
	message := responseBody["message"]

	if message != "subscription saved" {
		t.Fatalf("expected '%v', but received '%v'", "subscription saved", message)
	}
}

type MockMonogoService2 struct{}

func (ms *MockMonogoService2) Init(dbName string, connectionString string) {}
func (ms *MockMonogoService2) DB() *mongo.Database {
	return &mongo.Database{}
}
func (ms *MockMonogoService2) CreatedShardedCollection(collectionName string, field string, unique bool) {
}
func (ms *MockMonogoService2) CreateIndex(collectionName string, field string, unique bool) error {
	return nil
}
func (ms *MockMonogoService2) Find(collectionName string, filter primitive.D, items []interface{}) error {
	return nil
}
func (ms *MockMonogoService2) FindOne(collectionName string, filter primitive.D, item interface{}) error {
	return nil
}
func (ms *MockMonogoService2) InsertOne(collectionName string, item interface{}) error {
	return nil
}
func (ms *MockMonogoService2) UpdateOne(collectionName string, filter primitive.D, item interface{}, upsert bool) error {
	return errors.New("MOCK_ERROR")
}

func TestHandlerWithFaultyDB(t *testing.T) {
	mockMongoService := MockMonogoService2{}

	mockReq := models.Subscription{
		Company:  "fakepanda",
		UserID:   "abc@m.com",
		Endpoint: "https://fcm.googleapis.com/fcm/send/ePpws-p5uBI:APA91bFm_zyeVqFGxiw5kWJJR9KLy9eFRRUKbc8_yebfBbsNBcX3iJmAUgl22uu_zpe2Hv0RSHpEThblr_Bz5AbbHQ7WXuUO2LkxmqJjTc6N1GURuZEOjtt2Y6_pr0org506K1ZMk6lK",
		Keys: models.Keys{
			P256dh: "BKOynOSa_eUI1ZGhmWxsaA34lbfqtxGTXiZTFa24SpDjOQBHwwCfxLBuWRdO_92E5A3ia8VA3Q5774ECPK6-Khg",
			Auth:   "wquys90eGkBzfmTSyMT-PQ",
		},
	}

	objBytes, err := json.Marshal(mockReq)
	if err != nil {
		t.Fatal("Cannot marshal", err.Error())
	}
	request := httptest.NewRequest(http.MethodPost, "/api/subscriptions", bytes.NewBuffer(objBytes))

	writer := httptest.NewRecorder()
	handler(&mockMongoService, writer, request)
	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	responseBody := string(data)
	formats.Trace(responseBody)
	if !strings.Contains(responseBody, "MOCK_ERROR") {
		t.Fatalf("expected error to be '%s', but received '%v'", "MOCK_ERROR", responseBody)
	}
}
