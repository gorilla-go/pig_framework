package middlewares

import (
	"github.com/gorilla-go/pig"
	"io"
	"net/http"
	"pig_framework/config"
	"strconv"
)

type BodyLimit struct {
}

func (b *BodyLimit) Handle(context *pig.Context, f func(*pig.Context)) {
	length := context.Request().Raw().Header.Get("Content-Length")
	bytesCount := 0
	if length == "" {
		body := context.Request().Raw().Body
		defer func() {
			err := body.Close()
			if err != nil {
				panic(err)
			}
		}()
		bytes, err := io.ReadAll(body)
		if err != nil {
			panic(err)
		}
		bytesCount = len(bytes)
	} else {
		n, err := strconv.Atoi(length)
		if err != nil {
			panic(err)
		}
		bytesCount = n
	}

	maxMb := config.DefaultConfig[int]("REQUEST_BODY_LIMIT", 32)
	if bytesCount > maxMb*1024*1024 {
		context.Response().Raw().WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}
	f(context)
}
