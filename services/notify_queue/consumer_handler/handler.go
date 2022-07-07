package consumer_handler

import (
	"encoding/json"
	"net/http"

	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/logs"
	"github.com/web-notify/api/monorepo/services/notify_queue/models"
	// "github.com/web-notify/api/monorepo/services/notify_queue/models"
)

func Handler(rw http.ResponseWriter, req *http.Request) {
	logs.Trace("queue triggered")

	var requestBody qmodels.RequestBody
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// the message is stringified twice, so need to unmarshall twice
	var message string
	err = json.Unmarshal(requestBody.Data["req"], &message)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	logs.Trace("rawMessage:", message)
	err = json.Unmarshal([]byte(message), &message)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var subscription models.Subscription
	err = json.Unmarshal([]byte(message), &subscription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	logs.Trace("subscription:", subscription)

	responseBody := qmodels.ResponseBody{
		Outputs: map[string]interface{}{
			"res": "",
		},
		Logs:        []string{"Message successfully dequeued"},
		ReturnValue: "",
	}

	bytes, _ := json.Marshal(responseBody)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}
