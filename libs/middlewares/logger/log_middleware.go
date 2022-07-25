package logger

import (
	"net/http"

	"github.com/web-notify/api/monorepo/libs/middlewares"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
)

type logMiddleWare struct{}

func NewLogMiddleware() logMiddleWare {
	return logMiddleWare{}
}

func (mw logMiddleWare) Wrap(handler middlewares.Handler) middlewares.Handler {
	return func(rw http.ResponseWriter, req *http.Request) {
		// pre-process request here
		formats.Trace("LogMiddleware", "pre-processing request...")
		handler(rw, req)
		// post-process reponse here
	}
}
