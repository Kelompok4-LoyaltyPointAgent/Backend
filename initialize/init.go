package initialize

import (
	"github.com/kelompok4-loyaltypointagent/backend/cachedrepositories/cached_analytics_repository"
	"github.com/kelompok4-loyaltypointagent/backend/cachedrepositories/cached_invoiceurl_repository"
	"github.com/kelompok4-loyaltypointagent/backend/db"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/analytics_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/faq_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/favorites_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/feedback_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/forgot_password_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/hello_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/otp_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/product_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/transaction_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/user_handler"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/redisclient"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/analytics_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/faq_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/favorites_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/feedback_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/forgot_password_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/otp_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/packages_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_picture_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/transaction_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
	"github.com/kelompok4-loyaltypointagent/backend/services/analytics_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/faq_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/favorites_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/feedback_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/forgot_password_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/otp_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/product_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/transaction_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/user_service"
)

// Hello
var HelloHandler hello_handler.HelloHandler

// User
var userRepository user_repository.UserRepository
var userService user_service.UserService
var UserHandler user_handler.UserHandler

// OTP
var otpRepository otp_repository.OTPRepository
var otpService otp_service.OTPService
var OTPHandler otp_handler.OTPHandler

// Forgot Password
var forgotPasswordRepository forgot_password_repository.ForgotPasswordRepository
var forgotPasswordService forgot_password_service.ForgotPasswordService
var ForgotPasswordHandler forgot_password_handler.ForgotPasswordHandler

// Product
var productRepository product_repository.ProductRepository
var productService product_service.ProductService
var ProductHandler product_handler.ProductHandler

// Product Picture
var productPictureRepository product_picture_repository.ProductPictureRepository

// Credit
var creditRepository credit_repository.CreditRepository

// Packages
var packagesRepository packages_repository.PackagesRepository

// Transaction
var transactionRepository transaction_repository.TransactionRepository
var transactionService transaction_service.TransactionService
var TransactionHandler transaction_handler.TransactionHandler
var cachedInvoiceURLRepository cached_invoiceurl_repository.InvoiceURLRepository

// FAQ
var faqRepository faq_repository.FAQRepository
var faqService faq_service.FAQService
var FAQHandler faq_handler.FAQHandler

// Favorites
var favoritesRepository favorites_repository.FavoritesRepository
var favoritesService favorites_service.FavoritesService
var FavoritesHandler favorites_handler.FavoritesHandler

// Feedback
var FeedbackHandler feedback_handler.FeedbackHandler
var feedbackService feedback_service.FeedbackService
var feedbackRepository feedback_repository.FeedbackRepository

// Analytics
var AnalyticsHandler analytics_handler.AnalyticsHandler
var analyticsService analytics_service.AnalyticsService
var analyticsRepository analytics_repository.AnalyticsRepository
var cachedAnalyticsRepository cached_analytics_repository.CachedAnalyticsRepository

func Init() {
	helper.InitAppFirebase()
	initRepositories()
	initServices()
	initHandlers()
}

func initRepositories() {
	db := db.Init()

	userRepository = user_repository.NewUserRepository(db)
	productRepository = product_repository.NewProductRepository(db)
	creditRepository = credit_repository.NewCreditRepository(db)
	packagesRepository = packages_repository.NewPackagesRepository(db)
	productPictureRepository = product_picture_repository.NewProductPictureRepository(db)
	otpRepository = otp_repository.NewOTPRepository(db)
	transactionRepository = transaction_repository.NewTransactionRepository(db)
	faqRepository = faq_repository.NewFAQRepository(db)
	forgotPasswordRepository = forgot_password_repository.NewForgotPasswordRepository(db)
	favoritesRepository = favorites_repository.NewFavoritesRepository(db)
	feedbackRepository = feedback_repository.NewFeedbackRepository(db)
	analyticsRepository = analytics_repository.NewAnalyticsRepository(db)

	redisDB := redisclient.Init()
	cachedAnalyticsRepository = cached_analytics_repository.NewCachedAnalyticsRepository(redisDB)
	cachedInvoiceURLRepository = cached_invoiceurl_repository.NewInvoiceURLRepository(redisDB)
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
	productService = product_service.NewProductService(productRepository, creditRepository, packagesRepository, productPictureRepository)
	otpService = otp_service.NewOTPService(otpRepository, userRepository)
	transactionService = transaction_service.NewTransactionService(transactionRepository, productRepository, userRepository, cachedInvoiceURLRepository)
	faqService = faq_service.NewFAQService(faqRepository)
	forgotPasswordService = forgot_password_service.NewForgotPasswordService(forgotPasswordRepository, userRepository)
	favoritesService = favorites_service.NewFavoritesService(favoritesRepository)
	feedbackService = feedback_service.NewFeedbackService(feedbackRepository)
	analyticsService = analytics_service.NewAnalyticsService(analyticsRepository, cachedAnalyticsRepository)
}

func initHandlers() {
	HelloHandler = hello_handler.NewHelloHandler()
	UserHandler = user_handler.NewUserHandler(userService)
	ProductHandler = product_handler.NewProductHandler(productService)
	OTPHandler = otp_handler.NewOTPHandler(otpService)
	TransactionHandler = transaction_handler.NewTransactionHandler(transactionService)
	FAQHandler = faq_handler.NewFAQHandler(faqService)
	ForgotPasswordHandler = forgot_password_handler.NewForgotPasswordHandler(forgotPasswordService)
	FavoritesHandler = favorites_handler.NewFavoritesHandler(favoritesService)
	FeedbackHandler = feedback_handler.NewFeedbackHandler(feedbackService)
	AnalyticsHandler = analytics_handler.NewAnalyticsHandler(analyticsService)
}
