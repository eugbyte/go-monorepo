package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	hello "github.com/web-notify-lib/notify-api-azure/hello_handler"
	"github.com/web-notify-lib/notify-api-azure/lib/config"
)

func main() {
	var address string = config.Config.LOCAL_PORT
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		address = val
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", hello.Handler)

	log.Printf("About to listen on %s. Go to https://localhost:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), mux))

}
