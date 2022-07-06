package consumer_handler

import (
	"encoding/json"
	"net/http"

	qModels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/logs"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	var requestBody qModels.RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	logs.Trace("subscription", requestBody)

	responseBody := qModels.ResponseBody{}
	responseBody.Outputs.Response.Body = "success"
	response.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(responseBody)
	response.Write(bytes)
}
