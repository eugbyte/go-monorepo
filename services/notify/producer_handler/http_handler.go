package producer_handler

import (
	"context"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middlewares"
	vaultmwlib "github.com/web-notify/api/monorepo/libs/middlewares/vault"
	qlib "github.com/web-notify/api/monorepo/libs/queue"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/config"
)

// Dependency injection
func Handler(response http.ResponseWriter, request *http.Request) {
	queueName := "my-queue"
	var stage config.STAGE = config.Stage()
	queueAccountName := config.ENV_VARS[stage].QUEUE_ACCOUNT_NAME

	qBaseUrl := config.QueueBaseURL(stage, queueAccountName)
	var qService qlib.QueueServicer = qlib.NewQueueService(context.Background(), queueName, qBaseUrl, queueAccountName, config.ENV_VARS[stage].QUEUE_ACCOUNT_KEY)

	handler(qService, response, request)
}

var vaultService = vault.NewVaultService("https://kv-notify-secrets-stg.vault.azure.net")
var vaultMW = vaultmwlib.VaultMiddleware(vaultService)

// Apply middleware
var HTTPHandler http.Handler = middlewares.Middy(http.HandlerFunc(Handler), vaultMW)
