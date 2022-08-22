package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eugbyte/monorepo/libs/middleware"
	loggerMW "github.com/eugbyte/monorepo/libs/middleware/logger"
	"github.com/eugbyte/monorepo/services/greet/config"
	"github.com/eugbyte/monorepo/services/greet/hello_handler"
	"github.com/rs/cors"
)

func main() {
	var address string = config.New().LOCAL_PORT
	mux := http.NewServeMux()
	mux.Handle("/api/hello", hello_handler.HTTPHandler)

	wrappedMux := middleware.Middy(mux, loggerMW.LoggerMiddleware, cors.Default().Handler)

	log.Printf("STAGE: %s", config.Stage().String())
	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", address, address)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", address), wrappedMux))
}
