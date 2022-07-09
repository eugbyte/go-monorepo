package hello_handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type RequestBody struct {
	Message string `json:"message"`
}

func Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	message := strings.ToUpper("Hello") + "!!"
	responseBody := map[string]interface{}{"message": message}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
