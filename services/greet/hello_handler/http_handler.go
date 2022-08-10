package hello_handler

import "net/http"

var HTTPHandler http.Handler = http.HandlerFunc(Handler)
