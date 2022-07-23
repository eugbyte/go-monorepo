package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middlewares"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	hello "github.com/web-notify/api/monorepo/services/greet/hello_handler"
)

func main() {
	var stage config.STAGE = config.Stage()
	var address string = config.ENV_VARS[stage].LOCAL_PORT
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", middlewares.Middy(hello.Handler, middlewares.NewLogMiddleware()))

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), mux))
}
