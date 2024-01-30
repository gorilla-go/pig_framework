package middlewares

import "github.com/gorilla-go/pig"

type Auth struct {
}

func (a *Auth) Handle(context *pig.Context, f func(*pig.Context)) {
	f(context)
}
