package hello_handler

import (
	"encoding/json"
	"net/http"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	responseBody := map[string]interface{}{"message": "Hello World"}

	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(response).Encode(responseBody)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
