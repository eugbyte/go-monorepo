package producer_handler

import (
	"context"
	"encoding/json"
	"net/http"

	qLib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type RequestBody struct {
	Username string `json:"username"`
	Company  string `json:"company"`
}

func handler(qService qLib.QueueServicer, rw http.ResponseWriter, request *http.Request) {
	formats.Trace("In handler")

	var requestBody RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("requestBody", requestBody)

	message := formats.Stringify(requestBody)

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
	queueName := "my-queue"
	var stage config.STAGE = config.GetStage()
	queueAccountName := config.ENV_VARS[stage].QUEUE_ACCOUNT_NAME

	baseConnectionString := qLib.GetBaseConnectionString(stage, queueAccountName)
	var qService qLib.QueueServicer = qLib.NewQueueService(context.Background(), queueName, baseConnectionString, queueAccountName, config.ENV_VARS[stage].QUEUE_ACCOUNT_KEY)

	// Dependency injection
	handler(qService, response, request)
}
