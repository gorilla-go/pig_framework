package controller

import (
	"github.com/gorilla-go/pig"
)

func Index(ctx *pig.Context) {
	ctx.Response().Html(
		"<html><h2>P.I.G web framework</h2><a href=\"https://pig.gitbook.io/p.i.g-framework\">document</a></html>",
	)
}
