package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func RootPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("path load failed.")
	}
	return filepath.Dir(filename + "/../../")
}

func Config[T any](c string) T {
	se := os.Getenv(c)
	if se != "" {
		return any(se).(T)
	}
	return config[c].(T)
}

func DefaultConfig[T any](c string, def T) (ret T) {
	defer func() {
		if err := recover(); err != nil {
			ret = def
		}
	}()
	ret = Config[T](c)
	return
}
