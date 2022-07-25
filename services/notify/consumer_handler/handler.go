package consumer_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/db/mongo"
	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/lib"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/bson"
)

func handler(
	mongoService mongo.MonogoServicer,
	rw http.ResponseWriter,
	request *http.Request) {

	formats.Trace("queue triggered")

	var requestBody qmodels.RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var info models.Info
	info, err = lib.DecodeRawMassageToInfo(requestBody.Data["req"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("info:", info)

	id := fmt.Sprintf("%s__%s", info.Company, info.Username)
	var subscriber models.Subscription
	err = mongoService.FindOne("subscribers", bson.D{{Key: "_id", Value: id}}, &subscriber)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// TO DO
	// 1. Retrieve VAPID_PRIVATE_KEY from azure vault secret
	// 2. Web push with the private key

	qResponse := qmodels.ResponseBody{
		Outputs: map[string]interface{}{
			"res": "",
		},
		Logs:        []string{"Message successfully dequeued", fmt.Sprintf("message: '%s'", info)},
		ReturnValue: "",
	}

	objBytes, _ := json.Marshal(qResponse)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(objBytes)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	var stage config.STAGE = config.Stage()

	var mongoService mongo.MonogoServicer = mongo.NewMongoService("subscriberDB", config.ENV_VARS[stage].MONGO_DB_CONNECTION_STRING)
	// Dependency injection
	handler(mongoService, rw, req)
}
