package subscriber_handler

import (
	"net/http"

	mongolib "github.com/web-notify/api/monorepo/libs/db/mongo_lib"
	"github.com/web-notify/api/monorepo/libs/utils/config"
)

// Dependency injection
func Handler(rw http.ResponseWriter, request *http.Request) {
	var stage config.STAGE = config.Stage()
	var mongoService mongolib.MonogoServicer = mongolib.NewMongoService("subscriberDB", config.ENV_VARS[stage].MONGO_DB_CONNECTION_STRING)
	mongoService.CreatedShardedCollection(collectionName, "company", false)

	handler(mongoService, rw, request)
}

var HTTPHandler http.Handler = http.HandlerFunc(Handler)
