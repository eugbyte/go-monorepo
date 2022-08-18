package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eugbyte/monorepo/libs/formats"
	"github.com/eugbyte/monorepo/libs/middleware"
	"github.com/eugbyte/monorepo/services/webnotify/config"
	consumer "github.com/eugbyte/monorepo/services/webnotify/consumer_handler"
	producer "github.com/eugbyte/monorepo/services/webnotify/producer_handler"
	samplepush "github.com/eugbyte/monorepo/services/webnotify/sample_push_handler"
	subscribe "github.com/eugbyte/monorepo/services/webnotify/subscribe_handler"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var stage config.STAGE = config.Stage()
	formats.Trace("STAGE:", stage.String())
	var address string = config.New().LOCAL_PORT

	mux := http.NewServeMux()
	mux.Handle("/api/subscriptions", subscribe.HTTPHandler)
	mux.Handle("/api/notifications", producer.HTTPHandler)
	mux.Handle("/consumer_handler", consumer.HTTPHandler)
	mux.Handle("/api/sample-push", samplepush.HTTPHandler)

	wrappedMux := middleware.Middy(mux, cors.Default().Handler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), wrappedMux))
}
