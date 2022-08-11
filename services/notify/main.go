package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/web-notify/api/monorepo/libs/middlewares"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	consumer "github.com/web-notify/api/monorepo/services/notify/consumer_handler"
	producer "github.com/web-notify/api/monorepo/services/notify/producer_handler"
	samplepush "github.com/web-notify/api/monorepo/services/notify/sample_push_handler"
	subscribe "github.com/web-notify/api/monorepo/services/notify/subscribe_handler"
)

func main() {

	var stage config.STAGE = config.Stage()
	var address string = config.ENV_VARS[stage].LOCAL_PORT

	mux := http.NewServeMux()
	mux.Handle("/api/subscriptions", subscribe.HTTPHandler)
	mux.Handle("/api/notifications", producer.HTTPHandler)
	mux.Handle("/consumer_handler", consumer.HTTPHandler)
	mux.Handle("/api/sample-push", samplepush.HTTPHandler)

	wrappedMux := middlewares.Middy(mux, cors.Default().Handler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), wrappedMux))
}
