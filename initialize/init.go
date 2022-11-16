package initialize

import (
	userHandler "github.com/kelompok4-loyaltypointagent/backend/handler/user"
	userRepo "github.com/kelompok4-loyaltypointagent/backend/repositories/user"
	userService "github.com/kelompok4-loyaltypointagent/backend/services/user"
	"gorm.io/gorm"
)

//User
var user_repository userRepo.UserRepository
var user_service userService.UserService
var User_handler userHandler.UserHandler

func Init(db *gorm.DB) {
	initRepositories(db)
	initServices()
	initHandlers()
}

func initRepositories(db *gorm.DB) {
	user_repository = userRepo.NewUserRepository(db)
}

func initServices() {
	user_service = userService.NewUserService(user_repository)
}

func initHandlers() {
	User_handler = userHandler.NewUserHandler(user_service)
}
