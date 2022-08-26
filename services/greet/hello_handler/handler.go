package hello_handler

import (
	"encoding/json"
	"net/http"

	"github.com/eugbyte/monorepo/services/webnotify/config"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(rw, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	responseBody := map[string]string{
		"stage":   config.Stage().String(),
		"message": "Hello World",
	}

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
