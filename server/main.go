package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (hp *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return hp.templates.ExecuteTemplate(w, name, data)
}

func main() {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	e := echo.New()
	e.Static("css", "css")
	e.Static("assets", "assets")
	e.Renderer = renderer 

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home", "")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
