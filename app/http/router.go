package http

import (
	"github.com/gorilla-go/pig"
	"pig_framework/app/http/controller"
	"pig_framework/config"
)

func Router(r *pig.Router) *pig.Router {
	r.Static("/static/", config.RootPath()+"/static/")
	r.GET("/", controller.Index).Name("index")
	r.GET("/doc", func(context *pig.Context) {
		context.Response().Redirect("https://pig.gitbook.io/p.i.g-framework")
	})
	return r
}
