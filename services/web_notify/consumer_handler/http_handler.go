package consumerhandler

import (
	"log"
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/libs/middleware"
	webpush "github.com/eugbyte/monorepo/libs/notification/web_push"
	"github.com/eugbyte/monorepo/services/webnotify/config"
	"go.mongodb.org/mongo-driver/bson"
)

// Dependency injection

var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// Get the VAPID private key from azure key vault
	log.Println("queue trigger detected")

	privateKey := config.New().VAPID_PRIVATE_KEY
	publicKey := config.New().VAPID_PUBLIC_KEY
	email := config.New().VAPID_EMAIL

	log.Printf(formats.Stringify(bson.M{
		"vapidPublicKey":   publicKey,
		"vapidSenderEmail": email,
	}))

	webpushService := webpush.New(
		privateKey,
		publicKey,
		email,
	)
	mongoService := mongolib.New("subscriberDB", config.New().MONGO_DB_CONNECTION_STRING)

	log.Println("injecting services...")

	handler(webpushService, mongoService, rw, req)
})

// Wrap middlewares

var HTTPHandler http.Handler = middleware.Middy(httpHandler)
