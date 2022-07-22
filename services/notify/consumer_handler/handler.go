package consumer_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/web-notify/api/monorepo/libs/db/mongo"
	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/lib"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/bson"
)

func handler(
	mongoService mongo.MonogoServiceImp,
	vaultService vault.VaultServiceImpl,
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

	VAPID_PRIVATE_KEY, err := vaultService.GetSecret("VAPID_PRIVATE_KEY")
	if err != nil {
		http.Error(rw, errors.Wrap(err, "Unable to retrieve vapid private key from vault").Error(), http.StatusInternalServerError)
		return
	}
	formats.Trace(VAPID_PRIVATE_KEY)

	qResponse := qmodels.ResponseBody{
		Outputs: map[string]interface{}{
			"res": "",
		},
		Logs:        []string{"Message successfully dequeued", fmt.Sprintf("message: '%s'", info)},
		ReturnValue: "",
	}

	bytes, _ := json.Marshal(qResponse)
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
	var vaultService vault.VaultServiceImpl = &vault.VaultService{}
	vaultService.Init("abc")

	// Dependency injection
	handler(mongoService, vaultService, rw, req)
}
