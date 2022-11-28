package initialize

import (
	"github.com/kelompok4-loyaltypointagent/backend/db"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/credit_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/hello_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/product_handler"
	"github.com/kelompok4-loyaltypointagent/backend/handlers/user_handler"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/credit_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/packages_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_picture_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
	"github.com/kelompok4-loyaltypointagent/backend/services/credit_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/packages_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/product_service"
	"github.com/kelompok4-loyaltypointagent/backend/services/user_service"
)

// Hello
var HelloHandler hello_handler.HelloHandler

// User
var userRepository user_repository.UserRepository
var userService user_service.UserService
var UserHandler user_handler.UserHandler

// Product
var productRepository product_repository.ProductRepository
var productService product_service.ProductService
var ProductHandler product_handler.ProductHandler

//Product Picture
var productPictureRepository product_picture_repository.ProductPictureRepository

// Credit
var creditRepository credit_repository.CreditRepository
var creditService credit_service.CreditService
var CreditHandler credit_handler.CreditHandler

// Packages
var packagesRepository packages_repository.PackagesRepository
var packagesService packages_service.PackagesService

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
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
	productService = product_service.NewProductService(productRepository, creditRepository, packagesRepository, productPictureRepository)
	creditService = credit_service.NewCreditService(creditRepository)
	packagesService = packages_service.NewPackagesService(packagesRepository)
}

func initHandlers() {
	HelloHandler = hello_handler.NewHelloHandler()
	UserHandler = user_handler.NewUserHandler(userService)
	ProductHandler = product_handler.NewProductHandler(productService)
	CreditHandler = credit_handler.NewCreditHandler(creditService)
}
