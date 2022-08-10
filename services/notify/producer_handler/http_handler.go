package producer_handler

import (
	"context"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middlewares"
	"github.com/web-notify/api/monorepo/libs/middlewares/auth"
	qlib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/config"
)

// Dependency injection
func Handler(rw http.ResponseWriter, req *http.Request) {
	queueName := "my-queue"
	var stage config.STAGE = config.Stage()
	queueAccountName := config.ENV_VARS[stage].QUEUE_ACCOUNT_NAME

	qBaseUrl := config.QueueBaseURL(stage, queueAccountName)
	var qService qlib.QueueServicer = qlib.NewQueueService(context.Background(), queueName, qBaseUrl, queueAccountName, config.ENV_VARS[stage].QUEUE_ACCOUNT_KEY)

	handler(qService, rw, req)
}

var isAuth auth.IsAuth = func(header http.Header) (bool, error) {
	var vaultService = vault.NewVaultService("https://kv-notify-secrets-stg.vault.azure.net")
	key := header.Get("Notification-Key")
	company := header.Get("Company")
	checkVal, err := vaultService.GetSecret(company)
	if err != nil {
		return false, err
	}
	if key != checkVal {
		return false, nil
	}
	return true, nil
}

var HTTPHandler http.Handler = middlewares.Middy(http.HandlerFunc(Handler), auth.AuthMiddleware(isAuth))
