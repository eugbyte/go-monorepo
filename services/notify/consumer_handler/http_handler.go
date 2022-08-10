package consumer_handler

import (
	"net/http"

	mongolib "github.com/web-notify/api/monorepo/libs/db/mongo_lib"
	webpush "github.com/web-notify/api/monorepo/libs/notifications/web_push"
	appconfig "github.com/web-notify/api/monorepo/libs/store/app_config"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"github.com/web-notify/api/monorepo/services/notify/lib"
	"go.mongodb.org/mongo-driver/bson"
)

// Dependency injection
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

var HTTPHandler http.Handler = http.HandlerFunc(Handler)
