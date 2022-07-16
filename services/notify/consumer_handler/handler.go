package consumer_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func Handler(rw http.ResponseWriter, req *http.Request) {
	formats.Trace("queue triggered")

	var requestBody qmodels.RequestBody
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	formats.Trace(requestBody)

	// the message is stringified twice, so need to unmarshall twice
	var message string
	err = json.Unmarshal(requestBody.Data["req"], &message)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("rawMessage:", message)
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
	formats.Trace("subscription:", subscription)

	responseBody := qmodels.ResponseBody{
		Outputs: map[string]interface{}{
			"res": "",
		},
		Logs:        []string{"Message successfully dequeued", fmt.Sprintf("message: '%s'", message)},
		ReturnValue: "",
	}

	bytes, _ := json.Marshal(responseBody)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(bytes)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}