package subscriberhandler

import (
	"net/http"

	mongolib "github.com/eugbyte/monorepo/libs/db/mongo_lib"
	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/libs/middleware"
	"github.com/eugbyte/monorepo/services/webnotify/config"
)

// Dependency injection

var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	formats.Trace("in http subscribe handler")
	var mongoService mongolib.MonogoServicer = mongolib.New("subscriberDB", config.New().MONGO_DB_CONNECTION_STRING)
	handler(mongoService, rw, req)
})

// Wrap middlewares

var HTTPHandler http.Handler = middleware.Middy(httpHandler)
