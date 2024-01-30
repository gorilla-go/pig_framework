package helpers

import (
	"github.com/gorilla-go/pig"
	"html"
	"html/template"
	"os"
	"path/filepath"
	"pig_framework/config"
	"strings"
)

var templateCache = make(map[string][]byte)

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

	var file []byte
	var err error
	if config.Config[bool]("APP_DEBUG") == true {
		file, err = os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
	} else {
		if _, ok := templateCache[templatePath]; ok == false {
			file, err = os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			templateCache[templatePath] = file
		} else {
			file = templateCache[templatePath]
		}
	}

	tmpl, err := template.New(templatePath).
		Funcs(template.FuncMap(map[string]any{
			"render": render,
		})).
		Parse(string(file))
	if err != nil {
		panic(err)
	}
	w := &templateWriter{}
	err = tmpl.Execute(w, o)
	if err != nil {
		panic(err)
	}
	return html.UnescapeString(w.html)
}

func View(ctx *pig.Context, template string, o any, wrapper ...string) {
	if wrapper != nil && len(wrapper) > 0 {
		t := config.Config[string]("TEMPLATE_WRAPPER_STR")
		content := render(template, o)
		for _, wrapperPath := range wrapper {
			content = strings.Replace(
				render(wrapperPath, o),
				t,
				content,
				1,
			)
		}
		ctx.Response().Html(content)
		return
	}
	ctx.Response().Html(render(template, o))
}
