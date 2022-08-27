package consumerhandler

import (
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/libs/middleware"
	webpush "github.com/eugbyte/monorepo/libs/notification/web_push"
	"github.com/eugbyte/monorepo/libs/store/vault"
	"github.com/eugbyte/monorepo/services/webnotify/config"
)

// Dependency injection

var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// Get the VAPID private key from azure key vault
	formats.Trace("queue trigger detected")

	privateKey := config.New().VAPID_PRIVATE_KEY
	publicKey := config.New().VAPID_PUBLIC_KEY
	email := config.New().VAPID_EMAIL

	if config.Stage() == config.STAGING || config.Stage() == config.PROD {
		vaultService := vault.New(config.New().VAULT_URI)
		var fetchVal config.FetchVal = vaultService.GetSecret
		secrets, err := config.FetchAll(fetchVal, "vapid-private-key", "vapid-public-key", "vapid-email")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
		}
		privateKey = secrets[0]
		publicKey = secrets[1]
		email = secrets[2]
	}

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
