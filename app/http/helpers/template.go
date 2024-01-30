package helpers

import (
	"github.com/gorilla-go/pig"
	"html/template"
	"path/filepath"
	"pig_framework/config"
	"strings"
)

type templateWriter struct {
	html string
}

func (w *templateWriter) Write(p []byte) (n int, err error) {
	w.html += string(p)
	return len(p), nil
}

func render(templatePath string, o any) string {
	s := config.Config[string]("TEMPLATE_PATH")
	ext := config.Config[string]("TEMPLATE_EXT")
	filePath := filepath.Dir(s) + "/" + templatePath
	if ext != "" {
		filePath += "." + ext
	}

	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}
	w := &templateWriter{}
	err = tmpl.Execute(w, o)
	if err != nil {
		panic(err)
	}
	return w.html
}

func View(ctx *pig.Context, template string, o any, wrapper ...string) {
	if wrapper != nil && len(wrapper) > 0 {
		t := config.Config[string]("TEMPLATE_WRAPPER_STR")
		html := render(template, o)
		for _, wrapperPath := range wrapper {
			html = strings.Replace(
				render(wrapperPath, o),
				t,
				html,
				1,
			)
		}
		ctx.Response().Html(html)
		return
	}
	ctx.Response().Html(render(template, o))
}
