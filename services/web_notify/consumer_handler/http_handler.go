package consumerhandler

import (
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/libs/middleware"
	webpush "github.com/eugbyte/monorepo/libs/notification/web_push"
	appconfig "github.com/eugbyte/monorepo/libs/store/app_config"
	"github.com/eugbyte/monorepo/libs/store/vault"
	"github.com/eugbyte/monorepo/services/webnotify/config"
	"go.mongodb.org/mongo-driver/bson"
)

// Dependency injection

var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// Get the VAPID private key from azure key vault
	vaultService := vault.New(config.New().VAULT_URI)
	secrets, err := config.FetchAll(vaultService.GetSecret, "VAPID-PRIVATE-KEY")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	privateKey := secrets[0]

	// Get the VAPID public key and VAPID email from the app config
	appConfigService := appconfig.New("e53c986e-fa42-4065-bcef-9a5ae182d65a", "rg-webnotify-stg", "appcs-web-notify-stg-ea")
	var getConfig config.FetchVal = func(name string) (string, error) {
		return appConfigService.GetConfig(name, nil)
	}
	params, err := config.FetchAll(getConfig, "VAPID-PUBLIC-KEY", "VAPID-SENDER-EMAIL")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	publicKey := params[0]
	email := params[1]

	formats.Trace(bson.M{
		"vapidPublicKey":   publicKey,
		"vapidSenderEmail": email,
	})

	webpushService := webpush.New(
		privateKey,
		publicKey,
		email,
	)
	mongoService := mongolib.New("subscriberDB", config.New().MONGO_DB_CONNECTION_STRING)

	handler(webpushService, mongoService, rw, req)
})

// Wrap middlewares

var HTTPHandler http.Handler = middleware.Middy(httpHandler)