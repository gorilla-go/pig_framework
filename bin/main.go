package main

import (
	"github.com/gorilla-go/pig"
	"pig_framework/app/http"
	"pig_framework/config"
)

func main() {
	pig.New().
		Use(http.Middleware...).
		Pprof().
		Router(http.Router(pig.NewRouter())).
		Run(config.DefaultConfig("APP_PORT", 8080))
}
