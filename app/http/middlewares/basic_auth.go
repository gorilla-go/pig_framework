package middlewares

import (
	"github.com/gorilla-go/pig"
	"net/http"
	"pig_framework/config"
)

type BasicAuth struct {
}

func (a *BasicAuth) Handle(context *pig.Context, f func(*pig.Context)) {
	if !config.DefaultConfig[bool]("BASIC_AUTH_ENABLE", false) {
		f(context)
		return
	}

	username, password, ok := context.Request().Raw().BasicAuth()
	if !ok ||
		username != config.Config[string]("BASIC_AUTH_USERNAME") ||
		password != config.Config[string]("BASIC_AUTH_PASSWORD") {
		context.Response().Raw().Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		context.Response().Raw().WriteHeader(http.StatusUnauthorized)
		return
	}

	f(context)
}
