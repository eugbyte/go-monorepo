package producer_handler

import (
	"context"
	"encoding/json"
	"net/http"

	qLib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/logs"
	"github.com/web-notify/api/monorepo/services/notify_queue/models"
)

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

	if !(qService.QueueExist()) {
		logs.Trace("queue does not exist, creating one...")
		qService.CreateQueue(nil)
	} else {
		logs.Trace("queue exist")
	}

	_, err = qService.Enqueue(message, 0, 0)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	logs.Trace("successfully enqueued")

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(map[string]string{
		"message": "successfully enqueued",
	})
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(response http.ResponseWriter, request *http.Request) {
	var qService = qLib.QueueService{}
	queueName := "my-queue"
	qService.Init(context.Background(), queueName, config.QUEUE_ACCOUNT_NAME, config.QUEUE_ACCOUNT_KEY)

	// Dependency injection
	handler(&qService, response, request)
}
