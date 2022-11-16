package route

import (
	"github.com/kelompok4-loyaltypointagent/backend/handler"
	"github.com/kelompok4-loyaltypointagent/backend/initialize"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Setup(app *echo.Echo, db *gorm.DB) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	helloHandler := handler.NewHelloHandler()
	app.GET("/", helloHandler.Greeting)

	v1 := app.Group("/api/v1")
	v1.POST("/register", initialize.User_handler.CreateUser)
}
