package consumer_handler

import (
	"encoding/json"
	"net/http"
)

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var invokeRequest InvokeRequest

	d := json.NewDecoder(r.Body)
	d.Decode(&invokeRequest)

	outputs := make(map[string]interface{})
	outputs["message"] = ""

	resData := make(map[string]interface{})
	resData["body"] = "Order enqueued"
	outputs["res"] = resData
	invokeResponse := InvokeResponse{outputs, []string{"Hello log"}, nil}
	bytes, _ := json.Marshal(invokeResponse)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
