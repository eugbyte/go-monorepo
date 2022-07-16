package producer_handler

import (
	"context"
	"encoding/json"
	"net/http"

	qLib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func handler(qService qLib.QueueServiceImpl, rw http.ResponseWriter, request *http.Request) {
	formats.Trace("In handler")

	var subscription models.Subscription
	err := json.NewDecoder(request.Body).Decode(&subscription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("subscription", subscription)

	message := formats.Stringify(subscription)

	if !(qService.QueueExist()) {
		formats.Trace("queue does not exist, creating one...")
		err = qService.CreateQueue(nil)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		formats.Trace("queue exist")
	}

	_, err = qService.Enqueue(message, 0, 0)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	formats.Trace("successfully enqueued")

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(rw).Encode(map[string]string{
		"message": "successfully enqueued",
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(response http.ResponseWriter, request *http.Request) {
	var qService = qLib.QueueService{}
	queueName := "my-queue"
	rootConnection := qLib.GetConnectionString(config.STAGE, config.QUEUE_ACCOUNT_NAME)
	qService.Init(context.Background(), queueName, rootConnection, config.QUEUE_ACCOUNT_NAME, config.QUEUE_ACCOUNT_KEY)

	// Dependency injection
	handler(&qService, response, request)
}
