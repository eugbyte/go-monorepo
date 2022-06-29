package hello

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/web-notify-lib/notify-api-azure/src/lib/util"
)

type RequestBody struct {
	Message string `json:"message"`
}

func Handler(response http.ResponseWriter, request *http.Request) {

	var requestBody RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
	util.Trace("requestBody", requestBody)

	message := strings.ToUpper(requestBody.Message) + "!!"
	responseBody := map[string]interface{}{"message": message}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
