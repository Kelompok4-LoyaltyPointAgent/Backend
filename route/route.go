package route

import (
	"github.com/kelompok4-loyaltypointagent/backend/initialize"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", initialize.HelloHandler.Greeting)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.POST("/register", initialize.UserHandler.CreateUser)
	v1.POST("/login", initialize.UserHandler.Login)

}
