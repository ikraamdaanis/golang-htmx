package handler

import (
	"fmt"
	"net/http"

	"github.com/ikraamdaanis/golang-htmx/pkg/supabase"
	"github.com/labstack/echo/v4"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	loginBody := new(LoginBody)
	if err := c.Bind(loginBody); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	accessToken, err := supabase.LoginWithEmailAndPassword(loginBody.Email, loginBody.Password)

	if err != nil {
		fmt.Println("Error authenticating with Supabase: ", err)
		return err
	}

	return c.JSON(200, accessToken)
}
