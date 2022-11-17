package initialize

import (
	"github.com/kelompok4-loyaltypointagent/backend/db"
	user_handler "github.com/kelompok4-loyaltypointagent/backend/handler/user"
	user_repository "github.com/kelompok4-loyaltypointagent/backend/repositories/user"
	user_service "github.com/kelompok4-loyaltypointagent/backend/services/user"
)

// User
var userRepository user_repository.UserRepository
var userService user_service.UserService
var UserHandler user_handler.UserHandler

func Init() {
	initRepositories()
	initServices()
	initHandlers()
}

func initRepositories() {
	db := db.Init()

	userRepository = user_repository.NewUserRepository(db)
}

func initServices() {
	userService = user_service.NewUserService(userRepository)
}

func initHandlers() {
	UserHandler = user_handler.NewUserHandler(userService)
}
