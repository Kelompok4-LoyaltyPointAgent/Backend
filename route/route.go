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
	v1.POST("/transactions/webhook", initialize.TransactionHandler.TransactionWebhook)

	otp := v1.Group("/otp")
	otp.POST("/request", initialize.OTPHandler.RequestOTP)
	otp.POST("/verify", initialize.OTPHandler.VerifyOTP)

	forgotPassword := v1.Group("/forgot-password")
	forgotPassword.POST("/request", initialize.ForgotPasswordHandler.RequestForgotPassword)
	forgotPassword.POST("/submit", initialize.ForgotPasswordHandler.SubmitForgotPassword)

	users := v1.Group("/users", auth)
	//User
	users.GET("/me", initialize.UserHandler.FindUserByID)
	users.PUT("", initialize.UserHandler.UpdateUser)
	users.PUT("/change-password", initialize.UserHandler.ChangePassword)
	users.PUT("/reset-password", initialize.UserHandler.ChangePasswordFromResetPassword)
	users.POST("/check-password", initialize.UserHandler.CheckPassword)
	//Admin
	users.GET("", initialize.UserHandler.FindAllUser, middlewares.AuthorizedRoles([]string{"Admin"}))
	users.GET("/:id", initialize.UserHandler.FindUserByIDByAdmin, middlewares.AuthorizedRoles([]string{"Admin"}))
	users.PUT("/:id", initialize.UserHandler.UpdateUserByAdmin, middlewares.AuthorizedRoles([]string{"Admin"}))
	users.DELETE("/:id", initialize.UserHandler.DeleteUserByAdmin, middlewares.AuthorizedRoles([]string{"Admin"}))

	favorites := users.Group("/favorites", auth)
	favorites.GET("", initialize.FavoritesHandler.FindAll)
	favorites.POST("", initialize.FavoritesHandler.Create, middlewares.AuthorizedRoles([]string{"User"}))
	favorites.DELETE("/:product_id", initialize.FavoritesHandler.Delete, middlewares.AuthorizedRoles([]string{"User"}))

	faqs := v1.Group("/faqs", auth)
	faqs.GET("", initialize.FAQHandler.GetFAQs)
	faqs.POST("", initialize.FAQHandler.CreateFAQ, middlewares.AuthorizedRoles([]string{"Admin"}))
	faqs.GET("/:id", initialize.FAQHandler.GetFAQ)
	faqs.PUT("/:id", initialize.FAQHandler.UpdateFAQ, middlewares.AuthorizedRoles([]string{"Admin"}))
	faqs.DELETE("/:id", initialize.FAQHandler.DeleteFAQ, middlewares.AuthorizedRoles([]string{"Admin"}))

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

	transactionV1 := v1.Group("/transactions", auth)
	transactionV1.GET("", initialize.TransactionHandler.GetTransactions)
	transactionV1.POST("", initialize.TransactionHandler.CreateTransaction)
	transactionV1.GET("/:id", initialize.TransactionHandler.GetTransaction)
	transactionV1.PUT("/:id", initialize.TransactionHandler.UpdateTransaction, middlewares.AuthorizedRoles([]string{"Admin"}))
	transactionV1.DELETE("/:id", initialize.TransactionHandler.DeleteTransaction, middlewares.AuthorizedRoles([]string{"Admin"}))
	transactionV1.POST("/cancel/:id", initialize.TransactionHandler.CancelTransaction)

	feedback := v1.Group("/feedbacks", auth)
	feedback.GET("", initialize.FeedbackHandler.FindAll, middlewares.AuthorizedRoles([]string{"Admin"}))
	feedback.GET("/:id", initialize.FeedbackHandler.FindByID, middlewares.AuthorizedRoles([]string{"Admin"}))
	feedback.POST("", initialize.FeedbackHandler.Create)

	analytics := v1.Group("/analytics", auth)
	analytics.GET("", initialize.AnalyticsHandler.Analytics, middlewares.AuthorizedRoles([]string{"Admin"}))
}
