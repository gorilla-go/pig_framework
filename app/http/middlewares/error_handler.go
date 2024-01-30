package middlewares

import (
	"fmt"
	"github.com/gorilla-go/pig"
	"github.com/gorilla-go/pig/di"
	"github.com/sirupsen/logrus"
	"net/http"
	"pig_framework/config"
	"runtime/debug"
)

type ErrorHandlerService struct {
}

func (e *ErrorHandlerService) Handle(a any, context *pig.Context) {
	if config.DefaultConfig[bool]("APP_DEBUG", false) {
		context.Response().Text(fmt.Sprintf(
			"Error: %s\n\r%s",
			a,
			debug.Stack()),
		)
		return
	}

	logrus.WithFields(map[string]interface{}{}).Errorln(a)
	context.Response().Raw().WriteHeader(http.StatusInternalServerError)
}

type ErrorHandler struct {
}

func (e ErrorHandler) Handle(context *pig.Context, f func(*pig.Context)) {
	di.ProvideValue[pig.IHttpErrorHandler](context.Container(), &ErrorHandlerService{})
	f(context)
}
