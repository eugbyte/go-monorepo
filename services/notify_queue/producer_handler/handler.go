package producer_handler

import (
	"encoding/json"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/log"
	"github.com/web-notify/api/monorepo/services/notify_queue/models"
)

func Handler(response http.ResponseWriter, request *http.Request) {

	log.Trace("Init", "Hello")

	var subscription models.Subscription
	err := json.NewDecoder(request.Body).Decode(&subscription)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	log.Trace("subscription", subscription)

	responseBody := models.ResponseBody{}
	responseBody.Outputs.Message = log.Stringify(subscription)
	responseBody.Outputs.Response.Body = "Message received"

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	bytes, _ := json.Marshal(responseBody)
	response.Write(bytes)
}
