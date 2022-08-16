package consumer_handler

import (
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/libs/middleware"
	webpush "github.com/eugbyte/monorepo/libs/notification/web_push"
	appconfig "github.com/eugbyte/monorepo/libs/store/app_config"
	"github.com/eugbyte/monorepo/libs/store/vault"
	"github.com/eugbyte/monorepo/services/webnotify/config"
	"github.com/eugbyte/monorepo/services/webnotify/lib"
	"go.mongodb.org/mongo-driver/bson"
)

// Dependency injection
var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
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
		vapidConf.PrivateKey,
		vapidConf.PublicKey,
		vapidConf.Email,
	)
	mongoService := mongolib.NewMongoService("subscriberDB", config.ENV_VARS[stage].MONGO_DB_CONNECTION_STRING)

	handler(webpushService, mongoService, rw, req)
})

// Wrap middlewares
var HTTPHandler http.Handler = middleware.Middy(httpHandler)
