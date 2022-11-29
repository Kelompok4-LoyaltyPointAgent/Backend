package initialize

import (
	"github.com/kelompok4-loyaltypointagent/backend/db"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/hello_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/otp_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/product_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/user_handler"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/otp_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
	"github.com/kelompok4-loyaltypointagent/backend/services/otp_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/product_service"
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

// Credit
var creditRepository credit_repository.CreditRepository

func Init() {
	initRepositories()
	initServices()
	initHandlers()
}

func initRepositories() {
	db := db.Init()

	userRepository = user_repository.NewUserRepository(db)
	productRepository = product_repository.NewProductRepository(db)
	creditRepository = credit_repository.NewCreditRepository(db)
	otpRepository = otp_repository.NewOTPRepository(db)
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
	productService = product_service.NewProductService(productRepository, creditRepository)
	otpService = otp_service.NewOTPService(otpRepository, userRepository)
}

func initHandlers() {
	HelloHandler = hello_handler.NewHelloHandler()
	UserHandler = user_handler.NewUserHandler(userService)
	ProductHandler = product_handler.NewProductHandler(productService)
	OTPHandler = otp_handler.NewOTPHandler(otpService)
}
