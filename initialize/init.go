package initialize

import (
	"github.com/kelompok4-loyaltypointagent/backend/db"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/faq_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/hello_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/otp_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/product_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/transaction_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/user_handler"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/faq_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/otp_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/packages_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_picture_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/transaction_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
	"github.com/kelompok4-loyaltypointagent/backend/services/faq_service"
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

// FAQ
var faqRepository faq_repository.FAQRepository
var faqService faq_service.FAQService
var FAQHandler faq_handler.FAQHandler

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
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
	productService = product_service.NewProductService(productRepository, creditRepository, packagesRepository, productPictureRepository)
	otpService = otp_service.NewOTPService(otpRepository, userRepository)
	transactionService = transaction_service.NewTransactionService(transactionRepository, productRepository, userRepository)
	faqService = faq_service.NewFAQService(faqRepository)
}

func initHandlers() {
	HelloHandler = hello_handler.NewHelloHandler()
	UserHandler = user_handler.NewUserHandler(userService)
	ProductHandler = product_handler.NewProductHandler(productService)
	OTPHandler = otp_handler.NewOTPHandler(otpService)
	TransactionHandler = transaction_handler.NewTransactionHandler(transactionService)
	FAQHandler = faq_handler.NewFAQHandler(faqService)
}
