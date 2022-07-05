package middlewares

import "net/http"

type LogMiddleWare struct{}

func (mw *LogMiddleWare) Wrap(handler Handler) Handler {
	return func(response http.ResponseWriter, request *http.Request) {
		// pre-process request here
		handler(response, request)
		// post-process reponse here
	}
}
