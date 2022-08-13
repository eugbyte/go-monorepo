package subscriber_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/bson"
)

var collectionName = "subscribers"

func handler(mongoService mongolib.MonogoServicer, rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(rw, "Wrong HTTP Method", http.StatusBadRequest)
		return
	}

	var subscription models.Subscription
	err := json.NewDecoder(req.Body).Decode(&subscription)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace(collectionName, subscription)
	subscription.ID = fmt.Sprintf("%s__%s", subscription.Company, subscription.UserID)

	filter := bson.D{
		{Key: "_id", Value: subscription.ID},
		{Key: "company", Value: subscription.Company},
	}
	err = mongoService.UpdateOne(collectionName, filter, subscription, true)

	rw.Header().Set("Content-Type", "application/json")
	responseBody := make(map[string]string)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	} else {
		responseBody["message"] = "subscription saved"
	}

	err = json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
