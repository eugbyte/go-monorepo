package producer_handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eugbyte/monorepo/libs/middleware"
	"github.com/eugbyte/monorepo/libs/middleware/auth"
	qlib "github.com/eugbyte/monorepo/libs/queue"
	"github.com/eugbyte/monorepo/libs/store/vault"
	"github.com/eugbyte/monorepo/services/webnotify/config"
)

// Dependency injection
var httpHandler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	queueName := "my-queue"
	var stage config.STAGE = config.Stage()
	queueAccountName := config.New().QUEUE_ACCOUNT_NAME

	qBaseUrl := config.QueueBaseURL(stage, queueAccountName)
	var qService qlib.QueueServicer = qlib.New(context.Background(), queueName, qBaseUrl, queueAccountName, config.New().QUEUE_ACCOUNT_KEY)

	handler(qService, rw, req)
})

var isAuth auth.IsAuth = func(header http.Header) (bool, error) {
	company := header.Get("Notify-Secret-Name")
	key := header.Get("Notify-Secret-Value")

	var vaultService = vault.NewVaultService("https://kv-notify-secrets-stg.vault.azure.net")
	checkVal, err := vaultService.GetSecret(company)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if key != checkVal {
		return false, nil
	}
	return true, nil
}

// Wrap middlewares
var HTTPHandler http.Handler = middleware.Middy(httpHandler, auth.AuthMiddleware(isAuth))
