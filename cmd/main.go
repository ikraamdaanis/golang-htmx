package main

import (
	"log"
	"os"

	initializers "github.com/ikraamdaanis/golang-htmx/internal"
	"github.com/ikraamdaanis/golang-htmx/pkg/handler"
	types "github.com/ikraamdaanis/golang-htmx/pkg/models"
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
	e.GET("/", handler.Homepage)

	e.GET("/api/test", Test)
	e.POST("/api/auth/login", handler.Login)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func Test(c echo.Context) error {
	rows, _ := initializers.DB.Query("SELECT id, name, published, content, description, user_id, views, created_at, updated_at FROM forms")

	defer rows.Close()

	var formDataList []types.Form

	for rows.Next() {
		var formData types.Form
		err := rows.Scan(&formData.ID, &formData.Name, &formData.Published, &formData.Content, &formData.Description, &formData.UserId, &formData.Views, &formData.CreatedAt, &formData.UpdatedAt)
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
