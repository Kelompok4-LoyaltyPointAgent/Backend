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
	productV1.GET("/credits", initialize.ProductHandler.GetProductsWithCredits)
	productV1.POST("/credits", initialize.ProductHandler.CreateProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
	productV1.GET("/credits/:id", initialize.ProductHandler.GetProductWithCredit)
	productV1.PUT("/credits/:id", initialize.ProductHandler.UpdateProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
	productV1.DELETE("/credits/:id", initialize.ProductHandler.DeleteProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))

	productV1.GET("/package", initialize.ProductHandler.GetProductWithPackage)
	productV1.POST("/package", initialize.ProductHandler.CreateProductWithPackage, middlewares.AuthorizedRoles([]string{"Admin"}))
	productV1.GET("/package/:id", initialize.ProductHandler.GetProductWithPackage)
	productV1.PUT("/package/:id", initialize.ProductHandler.UpdateProductWithPackage, middlewares.AuthorizedRoles([]string{"Admin"}))
	productV1.DELETE("/package/:id", initialize.ProductHandler.DeleteProductWithPackage, middlewares.AuthorizedRoles([]string{"Admin"}))

}
