package main

import (
	"log"
	"net/http"

	"github.com/web-notify-lib/notify-api-azure/src/handlers/hello"
	"github.com/web-notify-lib/notify-api-azure/src/lib/config"
)

func main() {
	var addr string = config.Config.LOCAL_PORT
	log.Println(addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello.Handler)

	log.Printf("server is listening at %s...", addr)

	http.ListenAndServe(addr, mux)
}
