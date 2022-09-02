package graceecho

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type renderer struct {
	tpl *template.Template
}

func (r *renderer) Render(writer io.Writer, s string, i interface{}, context echo.Context) error {
	return r.tpl.ExecuteTemplate(writer, s, i)
}

func NewRenderer(tpl *template.Template) echo.Renderer {
	return &renderer{tpl: tpl}
}
