package middlewares

import (
	"context"
	"github.com/go-session/session"
	"github.com/gorilla-go/pig"
	"github.com/gorilla-go/pig/di"
	"pig_framework/config"
)

type Session struct {
	SessionStoreCtx context.Context
}

func (s *Session) Handle(ctx *pig.Context, f func(*pig.Context)) {
	if config.DefaultConfig[bool]("SESSION_ENABLE", true) == false {
		f(ctx)
		return
	}

	session.InitManager(
		session.SetCookieName(
			config.DefaultConfig[string]("SESSION_COOKIE_NAME", "p_sid"),
		),
		session.SetCookieLifeTime(
			config.DefaultConfig[int]("SESSION_COOKIE_LIFETIME", 3600),
		),
	)

	store, err := session.Start(
		s.SessionStoreCtx,
		ctx.Response().Raw(),
		ctx.Request().Raw(),
	)
	if err != nil {
		panic(err)
	}
	di.ProvideValue(ctx.Container(), store)
	f(ctx)
}
