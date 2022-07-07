package producer_handler

import (
	"context"
	"encoding/json"
	"net/http"

	qLib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/format"
	"github.com/web-notify/api/monorepo/services/notify_queue/models"
)

func handler(qService qLib.QueueServiceImpl, response http.ResponseWriter, request *http.Request) {
	format.Trace("In handler")

	var subscription models.Subscription
	err := json.NewDecoder(request.Body).Decode(&subscription)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	format.Trace("subscription", subscription)

	message := format.Stringify(subscription)

	if !(qService.QueueExist()) {
		format.Trace("queue does not exist, creating one...")
		qService.CreateQueue(nil)
	} else {
		format.Trace("queue exist")
	}

	_, err = qService.Enqueue(message, 0, 0)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	format.Trace("successfully enqueued")

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
