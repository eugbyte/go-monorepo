package hello_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/db/mongo"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
)

func handler(mongoService mongo.MonogoServiceImp, rw http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(rw, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	var subscription models.Subscription
	err := json.NewDecoder(request.Body).Decode(&subscription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("subscription", subscription)
	subscription.ID = fmt.Sprintf("%s__%s", subscription.Company, subscription.Username)

	mongoService.InsertOne("subscribers", subscription)

	responseBody := map[string]interface{}{"message": "subscription saved"}

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(rw http.ResponseWriter, request *http.Request) {
	var mongoService mongo.MonogoServiceImp = &mongo.MongoService{}
	mongoService.Init("subscriberDB", config.MONGO_DB_CONNECTION_STRING)
	mongoService.CreateIndex("subscribers", "company", false)
	handler(mongoService, rw, request)
}
