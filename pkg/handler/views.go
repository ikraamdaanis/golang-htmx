package handler

import (
	"github.com/ikraamdaanis/golang-htmx/pkg/utils"
	"github.com/labstack/echo/v4"
)

func Homepage(c echo.Context) error {
	return utils.RenderTemplate(c, "index", map[string]interface{}{"Title": "Home Page", "ButtonText": "Home Button"})
}
