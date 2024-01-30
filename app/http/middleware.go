package http

import (
	"context"
	"github.com/gorilla-go/pig"
	"pig_framework/app"
	"pig_framework/app/http/middlewares"
)

var Middleware = []pig.IMiddleware{
	&middlewares.ErrorHandler{},
	&middlewares.Cors{},
	&middlewares.BasicAuth{},
	&middlewares.Timezone{},
	&middlewares.RequestLogger{},
	&middlewares.Session{
		SessionStoreCtx: context.Background(),
	},
	&middlewares.BodyLimit{},
	&app.Provider{},
}
