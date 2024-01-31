package controller

import (
	"github.com/gorilla-go/pig"
	"pig_framework/app/http/helpers"
)

func Index(ctx *pig.Context) {
	helpers.View(ctx, "index", nil)
}
