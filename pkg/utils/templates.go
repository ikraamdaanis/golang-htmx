package utils

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

var Templates = template.Must(template.ParseGlob("views/*.html"))

func RenderTemplate(c echo.Context, tmpl string, data interface{}) error {
	err := Templates.ExecuteTemplate(c.Response().Writer, tmpl+".html", data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
