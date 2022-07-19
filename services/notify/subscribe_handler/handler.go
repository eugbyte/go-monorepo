package subscriber_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	mongoLib "github.com/web-notify/api/monorepo/libs/db/mongo"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "subscribers"

func handler(mongoService mongoLib.MonogoServiceImp, rw http.ResponseWriter, request *http.Request) {
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
	formats.Trace(collectionName, subscription)
	subscription.ID = fmt.Sprintf("%s__%s", subscription.Company, subscription.Username)

	responseBody := make(map[string]string)
	err = mongoService.InsertOne(collectionName, subscription)
	rw.Header().Set("Content-Type", "application/json")

	if mongo.IsDuplicateKeyError(err) {
		responseBody["message"] = "subscription already exists, skipping creation ..."
		rw.WriteHeader(http.StatusAccepted)
	} else {
		responseBody["message"] = "subscription saved"
	}

	err = json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Handler(rw http.ResponseWriter, request *http.Request) {
	var mongoService mongoLib.MonogoServiceImp = &mongoLib.MongoService{}
	mongoService.Init("subscriberDB", config.MONGO_DB_CONNECTION_STRING)
	mongoService.CreatedShardedCollection(collectionName, "company", false)

	// Dependency injection
	handler(mongoService, rw, request)
}
