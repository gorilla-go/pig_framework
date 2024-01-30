package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla-go/pig"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"pig_framework/config"
	"strings"
	"sync"
	"time"
)

var mutex sync.Mutex

type LoggerWriter struct {
}

func (l *LoggerWriter) Write(p []byte) (n int, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	fileName := fmt.Sprintf(
		"%saccess/%s/%s.log",
		config.Config[string]("LOG_PATH"),
		time.Now().Format("2006-01"),
		time.Now().Format("02"),
	)

	if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return 0, err
	}

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0755,
	)
	if err != nil {
		return 0, err
	}

	defer func() {
		e := file.Close()
		if err != nil {
			n = 0
			err = e
		}
	}()

	write, err := file.Write(p)
	if err != nil {
		return 0, err
	}
	return write, nil
}

type RequestLogger struct {
}

func (l *RequestLogger) Handle(context *pig.Context, f func(*pig.Context)) {
	logrus.SetOutput(&LoggerWriter{})
	if config.DefaultConfig[bool]("LOG_ENABLE", true) == false {
		f(context)
		return
	}

	startTime := time.Now()
	f(context)

	req := context.Request().Raw()

	headerMap := make(map[string]string)
	for s, stringArr := range req.Header {
		headerMap[s] = strings.Join(stringArr, ", ")
	}
	marshal, err := json.Marshal(headerMap)
	if err != nil {
		panic(err)
	}
	body := req.Body
	defer func() {
		err := body.Close()
		if err != nil {
			panic(err)
		}
	}()
	var content = make([]byte, 1024)
	n, err := body.Read(content)
	if err != nil && err != io.EOF {
		panic(err)
	}
	content = content[:n]
	logrus.WithFields(logrus.Fields{
		"method":   req.Method,
		"path":     req.URL.Path,
		"code":     context.Response().GetCode(),
		"timeUsed": fmt.Sprintf("%.2fms", float64(time.Since(startTime).Microseconds())/1000.0),
		"header":   string(marshal),
		"body":     string(content),
	}).Infoln("")
}
