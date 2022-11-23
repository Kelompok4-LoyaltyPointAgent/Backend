package route

import (
	"github.com/kelompok4-loyaltypointagent/backend/config"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/initialize"
	"github.com/kelompok4-loyaltypointagent/backend/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	authConfig := config.LoadAuthConfig()
	auth := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &helper.JWTCustomClaims{},
		SigningKey: []byte(authConfig.Secret),
		ContextKey: "token",
	})

	app.GET("/", initialize.HelloHandler.Greeting)

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.POST("/register", initialize.UserHandler.CreateUser)
	v1.POST("/login", initialize.UserHandler.Login)

	productV1 := v1.Group("/products", auth)
	productV1.GET("", initialize.ProductHandler.GetProducts)
	productV1.POST("", initialize.ProductHandler.CreateProduct, middlewares.AuthorizedRoles([]string{"Admin"}))

	creditV1 := v1.Group("/credits", auth)
	creditV1.GET("", initialize.CreditHandler.GetCredits)
	creditV1.POST("", initialize.CreditHandler.CreateCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
}
