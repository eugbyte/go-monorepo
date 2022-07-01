package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/api/web-notify-lib/services/notify_api/config"
	hello "github.com/api/web-notify-lib/services/notify_api/hello_handler"
)

func main() {
	var address string = config.Config.LOCAL_PORT

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", hello.Handler)

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), mux))

}
