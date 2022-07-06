package producer_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	qLib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/logs"
	"github.com/web-notify/api/monorepo/services/notify_queue/models"
)

var ctx = context.Background()

func handler(qService qLib.QueueServiceImpl, response http.ResponseWriter, request *http.Request) {

	logs.Trace("In handler")

	var subscription models.Subscription
	err := json.NewDecoder(request.Body).Decode(&subscription)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	logs.Trace("subscription", subscription)

	message := logs.Stringify(subscription)

	isQExist := qService.IsQueueExist()
	logs.Trace("isQExist:", isQExist)
	if !isQExist {
		logs.Trace("queue does not exist, creating one...")
		qService.CreateQueue()
	}

	messageUrl := qService.GetMessageURL()
	result, err := messageUrl.Enqueue(ctx, message, 0, 0)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(*result)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(response http.ResponseWriter, request *http.Request) {
	var qService qLib.QueueServiceImpl = &qLib.QueueService{}
	queueName := "my-queue"
	connectionString := fmt.Sprintf("%s/%s", "http://127.0.0.1:10001", "devstoreaccount1")
	logs.Trace("connectionString:", connectionString)
	qService.Init(queueName, ctx, config.QUEUE_ACCOUNT_NAME, config.QUEUE_ACCOUNT_KEY, connectionString)

	// Dependency injection
	handler(qService, response, request)
}
