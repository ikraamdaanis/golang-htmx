package initializers

import (
	"database/sql"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/ikraamdaanis/golang-htmx/pkg/utils"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error

	DB, err = sql.Open("postgres", os.Getenv("DB_URI"))

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	utils.PrintLine(color.FgGreen, color.Bold, color.Underline)("Successfully connected to the database.")
}
