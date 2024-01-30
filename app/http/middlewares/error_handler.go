package middleware

import (
	"fmt"
	"github.com/gorilla-go/pig"
	"github.com/gorilla-go/pig/di"
	"pig_framework/config"
	"runtime/debug"
)

type ErrorHandlerService struct {
}

func (e *ErrorHandlerService) Handle(a any, context *pig.Context) {
	d := config.DefaultConfig[bool]("app.debug", false)
	if d {
		context.Response().Text(fmt.Sprintf(
			"Error: %s\n%s",
			a,
			debug.Stack()),
		)
		return
	}

	context.Response().Raw().WriteHeader(500)
}

type ErrorHandler struct {
}

func (e ErrorHandler) Handle(context *pig.Context, f func(*pig.Context)) {
	di.ProvideValue[pig.IHttpErrorHandler](context.Container(), &ErrorHandlerService{})
	f(context)
}
