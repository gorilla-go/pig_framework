package controller

import (
	"github.com/gorilla-go/pig"
	"pig_framework/app/http/helpers"
	"time"
)

func Index(ctx *pig.Context) {
	helpers.View(ctx, "index", map[string]string{
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}
