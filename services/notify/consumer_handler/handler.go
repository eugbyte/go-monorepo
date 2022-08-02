package consumer_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	mongolib "github.com/web-notify/api/monorepo/libs/db/mongo_lib"
	webpush "github.com/web-notify/api/monorepo/libs/notifications/web_push"
	qmodels "github.com/web-notify/api/monorepo/libs/queue/models"
	appconfig "github.com/web-notify/api/monorepo/libs/store/app_config"
	"github.com/web-notify/api/monorepo/libs/store/vault"

	// appConfig "github.com/web-notify/api/monorepo/libs/store/app_config"

	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/lib"
	"github.com/web-notify/api/monorepo/services/notify/models"
	"go.mongodb.org/mongo-driver/bson"
)

func handler(
	webpushService webpush.WebPushServicer,
	mongoService mongolib.MonogoServicer,
	rw http.ResponseWriter,
	request *http.Request) {

	formats.Trace("queue triggered")

	var requestBody qmodels.RequestBody
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var info models.MessageInfo
	info, err = lib.DecodeRawMassageToInfo(requestBody.Data["req"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	formats.Trace("info:", info)

	id := fmt.Sprintf("%s__%s", info.Company, info.UserID)
	var subscriber models.Subscription
	err = mongoService.FindOne("subscribers", bson.D{{Key: "_id", Value: id}}, &subscriber)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	formats.Trace("sending notification...")
	err = webpushService.SendNotification(info.Notification, subscriber.Endpoint, subscriber.Keys.Auth, subscriber.Keys.P256dh, subscriber.ExpirationTime)
	if err != nil {
		formats.Trace(err.Error())
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	qResponse := qmodels.ResponseBody{
		Outputs: map[string]interface{}{
			"res": "",
		},
		Logs:        []string{"Message successfully dequeued", formats.Stringify(info)},
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

	// Get the VAPID keys
	vaultService := vault.NewVaultService("https://kv-notify-secrets-stg.vault.azure.net")
	appConfigService := appconfig.NewAppConfig("e53c986e-fa42-4065-bcef-9a5ae182d65a", "rg-webnotify-stg", "appcs-webnotify-stg")

	vapidConf, err := lib.FetchVapidConfig(vaultService, appConfigService)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	formats.Trace(bson.M{
		"vapidPublicKey":   vapidConf.PublicKey,
		"vapidSenderEmail": vapidConf.Email,
	})

	webpushService := webpush.NewWebPush(
		"olKYvd22Tt-OsP5-X25jxFfDAr4hlWFX6eeUX3i_D7I",
		"BPlL5OTZwtW-0-4pQXmobTgX6URszc9-UKoTTvpvInhUlPHorlDM8y04J-rrErlQXMVH7_Us983mNmmwsb-z53U",
		"eugenetham1994@gmail.com",
	)
	mongoService := mongolib.NewMongoService("subscriberDB", config.ENV_VARS[stage].MONGO_DB_CONNECTION_STRING)

	// Dependency injection
	handler(webpushService, mongoService, rw, req)
}
