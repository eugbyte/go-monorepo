package consumer_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/db/mongo"
	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Info struct {
	Username string `json:"username"`
	Company  string `json:"company"`
}

func handler(mongoService mongo.MonogoServiceImp, rw http.ResponseWriter, request *http.Request) {
	formats.Trace("queue triggered")

	var requestBody qmodels.RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	formats.Trace(requestBody)

	// the message is stringified twice, so need to unmarshall twice
	var message string
	err = json.Unmarshal(requestBody.Data["req"], &message)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("rawMessage:", message)
	err = json.Unmarshal([]byte(message), &message)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var info Info
	err = json.Unmarshal([]byte(message), &info)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("info:", info)

	id := fmt.Sprintf("%s__%s", info.Company, info.Username)
	var subscriber models.Subscription
	mongoService.FindOne("subscribers", bson.D{{Key: "_id", Value: id}}, &subscriber)

	responseBody := qmodels.ResponseBody{
		Outputs: map[string]interface{}{
			"res": "",
		},
		Logs:        []string{"Message successfully dequeued", fmt.Sprintf("message: '%s'", message)},
		ReturnValue: "",
	}

	bytes, _ := json.Marshal(responseBody)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(bytes)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	var mongoService mongo.MonogoServiceImp = &mongo.MongoService{}
	mongoService.Init("subscriberDB", config.MONGO_DB_CONNECTION_STRING)
	handler(mongoService, rw, req)
}
