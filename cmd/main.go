package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.Static("/", "assets")
	e.GET("/", Homepage)
	e.GET("/hello", Hello)

	e.Logger.Fatal(e.Start(":5000"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Homepage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "World")
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}
