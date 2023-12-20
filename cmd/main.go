package main

import (
	views "github.com/ikraamdaanis/golang-htmx/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "public")
	e.GET("/", Homepage)
	e.GET("/hello", Hello)

	e.Logger.Fatal(e.Start(":5001"))
}

func Homepage(c echo.Context) error {
	return views.Index().Render(c.Request().Context(), c.Response())
}

func Hello(c echo.Context) error {
	return views.Test("Ikraam").Render(c.Request().Context(), c.Response())
}
