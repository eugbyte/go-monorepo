package consumer_handler

import (
	"encoding/json"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/logs"
)

func Handler(rw http.ResponseWriter, req *http.Request) {
	var requestBody models.RequestBody

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	logs.Trace("requestBody:", requestBody)

	var message string
	err = json.Unmarshal(requestBody.Data["req"], &message)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	logs.Trace("message:", message)

	outputs := models.Output{}
	outputs.Message = "some message"
	outputs.Res.Body = message
	responseBody := models.ResponseBody{
		Outputs:     outputs,
		Logs:        []string{"Message successfully dequeued"},
		ReturnValue: message,
	}

	bytes, _ := json.Marshal(responseBody)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}
