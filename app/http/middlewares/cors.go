package middlewares

import (
	"github.com/gorilla-go/pig"
	"net/http"
)

type Cors struct {
}

func (c Cors) Handle(context *pig.Context, f func(*pig.Context)) {
	if context.Request().IsOption() {
		context.Response().Raw().Header().Set("Access-Control-Allow-Origin", "*")
		context.Response().Raw().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Response().Raw().Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		context.Response().Raw().WriteHeader(http.StatusNoContent)
		return
	}
	f(context)
}
