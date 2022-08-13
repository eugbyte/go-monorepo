package subscriber_handler

import (
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/middleware"
	"github.com/eugbyte/monorepo/services/web-push/config"
)

// Dependency injection
var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	var stage config.STAGE = config.Stage()
	var mongoService mongolib.MonogoServicer = mongolib.NewMongoService("subscriberDB", config.ENV_VARS[stage].MONGO_DB_CONNECTION_STRING)
	mongoService.CreatedShardedCollection(collectionName, "company", false)

	handler(mongoService, rw, req)
})

// Wrap middlewares
var HTTPHandler http.Handler = middleware.Middy(httpHandler)
