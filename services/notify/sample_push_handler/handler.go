package sample_push_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(rw, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	var info models.MessageInfo
	err := json.NewDecoder(req.Body).Decode(&info)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace(info)

	objBytes, err := json.Marshal(info)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// make a request to the producer_handler
	// only way to invoke one lambda from another, is through Durable Functions, which is not available in golang
	// https://docs.microsoft.com/en-us/azure/azure-functions/durable/durable-functions-overview?tabs=csharp#language-support
	client := &http.Client{}
	post, err := http.NewRequest("POST", config.ENV_VARS[config.DEV].NOTIFY_PRODUCER_URL, bytes.NewBuffer(objBytes))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Header.Set("API-Key", "sample-api-key")
	post.Header.Set("Notify-Secret-Name", "demo-company")
	post.Header.Set("Notify-Secret-Value", "YfZUV8HgaA4tMuH")

	formats.Trace("sending request...")
	resp, err := client.Do(post)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Sprintln(string(body))

	err = json.NewEncoder(rw).Encode(map[string]string{
		"message": "successfully enqueued",
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}