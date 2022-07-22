package middlewares

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type LogMiddleWare struct{}

func NewLogMiddleware() LogMiddleWare {
	return LogMiddleWare{}
}

func (mw LogMiddleWare) Wrap(handler Handler) Handler {
	return func(response http.ResponseWriter, request *http.Request) {
		// pre-process request here
		formats.Trace("LogMiddleware", "pre-processing request...")
		handler(response, request)
		// post-process reponse here
	}
}
