package producer_handler

import (
	"encoding/json"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/log"
	"github.com/web-notify/api/monorepo/services/notify_queue/models"
)

type RequestBody struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type QueueData struct {
	Url    string
	Method string
	Query  map[string]string

	Headers struct {
		ContentType []string `json:"Content-Type"`
	}
	Params map[string]string
	Body   interface{}
}

type Outputs struct {
	Message  string `json:"message"`
	Response struct {
		Body string `json:"body"`
	} `json:"response"`
}

type ResponseBody struct {
	Outputs     Outputs
	Logs        []string
	ReturnValue interface{}
}

func Handler(response http.ResponseWriter, request *http.Request) {

	var requestBody RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	log.Trace("requestBody.MetaData", requestBody.Metadata)

	var reqData QueueData
	json.Unmarshal(requestBody.Data["request"], &reqData)
	log.Trace("reqData", reqData)

	var subscription models.Subscription
	bytes, _ := json.Marshal(reqData.Body)
	err = json.Unmarshal(bytes, &subscription)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody := ResponseBody{}
	responseBody.Outputs.Message = log.Stringify(subscription)
	responseBody.Outputs.Response.Body = "Message received"

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")

	bytes, _ = json.Marshal(responseBody)
	response.Write(bytes)
}
