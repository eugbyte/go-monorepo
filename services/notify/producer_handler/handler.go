package producer_handler

import (
	"encoding/json"
	"net/http"

	"github.com/eugbyte/monorepo/libs/formats"
	qlib "github.com/eugbyte/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func handler(qService qlib.QueueServicer, rw http.ResponseWriter, request *http.Request) {
	formats.Trace("In handler")

	var info models.MessageInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("requestBody", info)

	message := formats.Stringify(info)

	if !(qService.QueueExist()) {
		formats.Trace("queue does not exist, creating one...")
		err = qService.CreateQueue(nil)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		formats.Trace("queue exist")
	}

	_, err = qService.Enqueue(message, 0, 0)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	formats.Trace("successfully enqueued")

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(rw).Encode(map[string]string{
		"message": "successfully enqueued",
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
