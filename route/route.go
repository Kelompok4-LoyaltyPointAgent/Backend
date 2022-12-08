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
	app.Use(middleware.CORS())

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

	otp := v1.Group("/otp")
	otp.POST("/request", initialize.OTPHandler.RequestOTP)
	otp.POST("/verify", initialize.OTPHandler.VerifyOTP)

	users := v1.Group("/users", auth)
	//User
	users.GET("/me", initialize.UserHandler.FindUserByID)
	users.PUT("", initialize.UserHandler.UpdateUser)
	users.PUT("/change-password", initialize.UserHandler.ChangePassword)
	//Admin
	users.GET("", initialize.UserHandler.FindAllUser, middlewares.AuthorizedRoles([]string{"Admin"}))
	users.GET("/:id", initialize.UserHandler.FindUserByIDByAdmin, middlewares.AuthorizedRoles([]string{"Admin"}))
	users.PUT("/:id", initialize.UserHandler.UpdateUserByAdmin, middlewares.AuthorizedRoles([]string{"Admin"}))
	users.DELETE("/:id", initialize.UserHandler.DeleteUserByAdmin, middlewares.AuthorizedRoles([]string{"Admin"}))

	faqs := v1.Group("/faqs", auth)
	faqs.GET("", initialize.ProductHandler.GetProductsWithCredits)
	faqs.POST("", initialize.ProductHandler.CreateProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
	faqs.GET("/:id", initialize.ProductHandler.GetProductWithCredit)
	faqs.PUT("/:id", initialize.ProductHandler.UpdateProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
	faqs.DELETE("/:id", initialize.ProductHandler.DeleteProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))

	products := v1.Group("/products")

	credits := products.Group("/credits", auth)
	credits.GET("", initialize.ProductHandler.GetProductsWithCredits)
	credits.POST("", initialize.ProductHandler.CreateProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
	credits.GET("/:id", initialize.ProductHandler.GetProductWithCredit)
	credits.PUT("/:id", initialize.ProductHandler.UpdateProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))
	credits.DELETE("/:id", initialize.ProductHandler.DeleteProductWithCredit, middlewares.AuthorizedRoles([]string{"Admin"}))

	packages := products.Group("/packages", auth)
	packages.GET("", initialize.ProductHandler.GetProductsWithPackages)
	packages.POST("", initialize.ProductHandler.CreateProductWithPackage, middlewares.AuthorizedRoles([]string{"Admin"}))
	packages.GET("/:id", initialize.ProductHandler.GetProductWithPackage)
	packages.PUT("/:id", initialize.ProductHandler.UpdateProductWithPackage, middlewares.AuthorizedRoles([]string{"Admin"}))
	packages.DELETE("/:id", initialize.ProductHandler.DeleteProductWithPackage, middlewares.AuthorizedRoles([]string{"Admin"}))

}
