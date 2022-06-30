package main

import (
	"log"
	"net/http"

	"github.com/web-notify-lib/notify-api-azure/src/handlers/hello"
	"github.com/web-notify-lib/notify-api-azure/src/lib/config"
)

func main() {
	var address string = config.Config.LOCAL_PORT
	log.Println(address)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", hello.Handler)

	log.Printf("server is listening at %s...", address)
	log.Fatal(http.ListenAndServe(address, mux))

}
