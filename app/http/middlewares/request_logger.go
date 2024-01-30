package middlewares

import "github.com/gorilla-go/pig"

type Logger struct {
}

func (l Logger) Handle(context *pig.Context, f func(*pig.Context)) {
	f(context)
}
