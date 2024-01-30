package http

import (
	"github.com/gorilla-go/pig"
	"pig_framework/app/http/controller"
)

func Router(r *pig.Router) *pig.Router {
	r.GET("/", controller.Index)
	r.GET("/doc", func(context *pig.Context) {
		context.Response().Redirect("https://pig.gitbook.io/p.i.g-framework")
	})
	return r
}
