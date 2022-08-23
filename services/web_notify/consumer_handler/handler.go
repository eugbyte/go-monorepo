package consumerhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	webpush "github.com/eugbyte/monorepo/libs/notification/web_push"
	qmodels "github.com/eugbyte/monorepo/libs/queue/models"

	// appConfig "github.com/web-notify/api/monorepo/libs/store/app_config"

	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/services/webnotify/models"
	"go.mongodb.org/mongo-driver/bson"
)

func handler(
	webpushService webpush.WebPushServicer,
	mongoService mongolib.MonogoServicer,
	rw http.ResponseWriter,
	req *http.Request) {

	log.Println("in handler.go")

	var requestBody qmodels.RequestBody
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("requestBody", formats.Stringify(requestBody))

	var info models.MessageInfo
	info, err = decodeRawMassageToInfo(requestBody.Data["req"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("info:", info)

	id := fmt.Sprintf("%s__%s", info.Company, info.UserID)
	var subscriber models.Subscription
	err = mongoService.FindOne("subscribers", bson.D{{Key: "_id", Value: id}}, &subscriber)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("sending notification...")
	err = webpushService.SendNotification(info.Notification, subscriber.Endpoint, subscriber.Keys.Auth, subscriber.Keys.P256dh, subscriber.ExpirationTime)
	if err != nil {
		log.Println(err.Error())
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
	log.Println("Message successfully dequeued...")
}
