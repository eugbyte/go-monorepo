package middlewares

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/logs"
)

type LogMiddleWare struct{}

func (mw LogMiddleWare) Wrap(handler Handler) Handler {
	return func(response http.ResponseWriter, request *http.Request) {
		// pre-process request here
		logs.Trace("LogMiddleware", "pre-processing request...")
		handler(response, request)
		// post-process reponse here
	}
}
