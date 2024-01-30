package middlewares

import (
	"github.com/gorilla-go/pig"
	"pig_framework/config"
	"time"
)

type Timezone struct {
}

func (t Timezone) Handle(context *pig.Context, f func(*pig.Context)) {
	location, err := time.LoadLocation(
		config.DefaultConfig("DEFAULT_TIMEZONE", "Asia/Shanghai"),
	)
	if err != nil {
		panic(err)
	}
	time.Local = location
	f(context)
}
