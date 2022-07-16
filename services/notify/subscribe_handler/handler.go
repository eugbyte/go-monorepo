package hello_handler

import (
	"encoding/json"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	var subscription models.Subscription
	err := json.NewDecoder(request.Body).Decode(&subscription)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("subscription", subscription)

	responseBody := map[string]interface{}{"message": "subscription saved"}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
