package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/web-notify/api/monorepo/libs/middlewares"
	loggerMW "github.com/web-notify/api/monorepo/libs/middlewares/logger"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	hello "github.com/web-notify/api/monorepo/services/greet/hello_handler"
)

func main() {
	var stage config.STAGE = config.Stage()
	var address string = config.ENV_VARS[stage].LOCAL_PORT
	mux := http.NewServeMux()
	mux.Handle("/api/hello", middlewares.Middy(http.HandlerFunc(hello.Handler), loggerMW.LoggerMiddleware))

	wrappedMux := cors.Default().Handler(mux)

	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), wrappedMux))
}
