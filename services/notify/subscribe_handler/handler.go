package hello_handler

import (
	"encoding/json"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func Handler(response http.ResponseWriter, request *http.Request) {
<<<<<<< HEAD
	// (response).Header().Set("Access-Control-Allow-Origin", "*")
	// (response).Header().Set("Access-Control-Allow-Origin", "*")
	// (response).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (response).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// if request.Method == http.MethodOptions {
	// 	return
	// }

=======
>>>>>>> 33398c7 (chore: rename notify_queue to notify)
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
<<<<<<< HEAD
	// response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// response.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type")
=======
>>>>>>> 33398c7 (chore: rename notify_queue to notify)
	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
