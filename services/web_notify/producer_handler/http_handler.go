package producerhandler

import (
	"context"
	"net/http"

	"github.com/eugbyte/monorepo/libs/middleware"
	qlib "github.com/eugbyte/monorepo/libs/queue"
	"github.com/eugbyte/monorepo/services/webnotify/config"
)

// Dependency injection

var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	var stage config.STAGE = config.Stage()
	queueName := "stq-webnotify"
	queueAccountName := config.New().QUEUE_ACCOUNT_NAME

	qBaseUrl := config.QueueBaseURL(stage, queueAccountName)
	var qService qlib.QueueServicer = qlib.New(context.Background(), queueName, qBaseUrl, queueAccountName, config.New().QUEUE_ACCOUNT_KEY)

	handler(qService, rw, req)
})

// Wrap middlewares

var HTTPHandler http.Handler = middleware.Middy(httpHandler)
