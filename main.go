package main

import (
	"fmt"
	"log"
	"net/http"

	hello "github.com/web-notify-lib/notify-api-azure/hello_handler"
	"github.com/web-notify-lib/notify-api-azure/lib/config"
)

func main() {
	var address string = config.Config.LOCAL_PORT

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", hello.Handler)

	fmt.Printf("server is listening at %s...", address)
	log.Fatal(http.ListenAndServe(address, mux))

}
