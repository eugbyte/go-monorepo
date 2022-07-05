package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/config"
	producer "github.com/web-notify/api/monorepo/services/notify_queue/producer_handler"
)

func main() {
	var address string = config.Config.LOCAL_PORT
	mux := http.NewServeMux()
	mux.HandleFunc("/api/notify", producer.Handler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), mux))
}
