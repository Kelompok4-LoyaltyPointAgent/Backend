package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelompok4-loyaltypointagent/backend/config"
	"github.com/kelompok4-loyaltypointagent/backend/initialize"
	"github.com/kelompok4-loyaltypointagent/backend/route"
	"github.com/labstack/echo/v4"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	initialize.Init()

	app := echo.New()
	route.Setup(app)

	httpConfig := config.LoadHTTPConfig()
	app.Logger.Fatal(app.Start(":" + httpConfig.Port))
}
