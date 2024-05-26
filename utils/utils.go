package utils

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

func ToLog(obj interface{}) string {
	data, _ := json.Marshal(&obj)

	return string(data)
}

type Template struct {
	Templates *template.Template
}

var Temps = &Template{
	Templates: template.Must(template.ParseGlob("views/*.go.html")),
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
