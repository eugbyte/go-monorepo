package samplepushhandler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/services/webnotify/config"
	"github.com/eugbyte/monorepo/services/webnotify/models"
)

func handler(client *http.Client, rw http.ResponseWriter, req *http.Request) {
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
	post, err := http.NewRequest("POST", config.New().NOTIFY_BASE_URL+"/notifications", bytes.NewBuffer(objBytes))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Header.Set("x-functions-key", config.EnvOrDefault("X_FUNCTIONS_KEY_SUBSCRIBE_HANDLER_DEMO_COMPANY", ""))

	formats.Trace("sending request...")
	resp, err := client.Do(post)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
