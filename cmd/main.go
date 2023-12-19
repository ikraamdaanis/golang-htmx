package main

import (
	"html/template"
	"io"
	"net/http"

	views "github.com/ikraamdaanis/golang-htmx/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t

	e.Static("/", "public")
	e.GET("/", Homepage)
	e.GET("/hello", Hello)

	e.Logger.Fatal(e.Start(":5001"))
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
	// data := map[string]interface{}{
	// 	"message": "This is a message",
	// }

	component := views.Test("Ikraam").Render(c.Request().Context(), c.Response())

	return component
}
