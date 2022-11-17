package route

import (
	"github.com/kelompok4-loyaltypointagent/backend/handler"
	"github.com/kelompok4-loyaltypointagent/backend/initialize"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	helloHandler := handler.NewHelloHandler()
	app.GET("/", helloHandler.Greeting)

	v1 := app.Group("/v1")
	v1.POST("/register", initialize.UserHandler.CreateUser)
}
