package producer_handler

import (
	"context"
	"encoding/json"
	"net/http"

	qlib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func handler(qService qlib.QueueServicer, rw http.ResponseWriter, request *http.Request) {
	formats.Trace("In handler")

	var info models.MessageInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("requestBody", info)

	message := formats.Stringify(info)

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

// Dependency injection
func Handler(response http.ResponseWriter, request *http.Request) {
	queueName := "my-queue"
	var stage config.STAGE = config.Stage()
	queueAccountName := config.ENV_VARS[stage].QUEUE_ACCOUNT_NAME

	qBaseUrl := config.QueueBaseURL(stage, queueAccountName)
	var qService qlib.QueueServicer = qlib.NewQueueService(context.Background(), queueName, qBaseUrl, queueAccountName, config.ENV_VARS[stage].QUEUE_ACCOUNT_KEY)

	handler(qService, response, request)
}

// var vaultService = vault.NewVaultService("https://kv-notify-secrets-stg.vault.azure.net")

// Apply middleware
var HTTPHandler http.Handler = http.HandlerFunc(Handler)
