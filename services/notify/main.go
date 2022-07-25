package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/web-notify/api/monorepo/libs/middlewares"
	vaultMWLib "github.com/web-notify/api/monorepo/libs/middlewares/vault"
	"github.com/web-notify/api/monorepo/libs/store/vault"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	consumer "github.com/web-notify/api/monorepo/services/notify/consumer_handler"
	producer "github.com/web-notify/api/monorepo/services/notify/producer_handler"
	subscribe "github.com/web-notify/api/monorepo/services/notify/subscribe_handler"
)

func main() {

	var vaultService vault.VaultServicer
	vaultService = vault.NewVaultService("https://kv-customers-stg.vault.azure.net/")
	vaultMW := vaultMWLib.VaultMiddleware(vaultService)

	var stage config.STAGE = config.Stage()
	var address string = config.ENV_VARS[stage].LOCAL_PORT

	mux := http.NewServeMux()
	mux.Handle("/api/notifications", middlewares.Middy(http.HandlerFunc(producer.Handler), vaultMW))
	mux.HandleFunc("/consumer_handler", consumer.Handler)
	mux.HandleFunc("/api/subscriptions", subscribe.Handler)

	wrappedMux := middlewares.Middy(mux, cors.Default().Handler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), wrappedMux))
}
