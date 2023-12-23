package initializers

import (
	"log"

	"github.com/fatih/color"
	"github.com/ikraamdaanis/golang-htmx/pkg/utils"
	"github.com/joho/godotenv"
)

func LoadSecretKeys() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("There was an error loading .env file.")
	}

	utils.PrintLine(color.FgCyan)("Successfully loaded the .env file variables.")
}
