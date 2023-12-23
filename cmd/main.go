package main

import (
	"log"
	"os"

	initializers "github.com/ikraamdaanis/golang-htmx/internal"
	views "github.com/ikraamdaanis/golang-htmx/views"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func init() {
	initializers.LoadSecretKeys()
	initializers.ConnectDatabase()
}

func main() {
	e := echo.New()

	e.Static("/", "public")
	e.GET("/", Homepage)
	e.GET("/hello", Hello)
	e.GET("/api/test", Test)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func Homepage(c echo.Context) error {
	return views.Index().Render(c.Request().Context(), c.Response())
}

func Hello(c echo.Context) error {
	return views.Test("Ikraam").Render(c.Request().Context(), c.Response())
}

type FormData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Add more fields based on your database schema
}

func Test(c echo.Context) error {
	rows, _ := initializers.DB.Query("SELECT id, name FROM forms")

	defer rows.Close()

	var formDataList []FormData

	for rows.Next() {
		var formData FormData
		err := rows.Scan(&formData.ID, &formData.Name)
		if err != nil {
			log.Fatal(err)
		}
		formDataList = append(formDataList, formData)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return c.JSON(200, formDataList)
}
