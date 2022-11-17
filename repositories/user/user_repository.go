package user

import (
	"github.com/kelompok4-loyaltypointagent/backend/models/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id string) (user.User, error)
	Create(user user.User) (user.User, error)
	FindAll() ([]user.User, error)
	Update(user user.User, id string) (user.User, error)
	Delete(id string) (user.User, error)
	FindByEmail(email string) (user.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindByID(id string) (user.User, error) {
	var user user.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) Create(user user.User) (user.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) FindAll() ([]user.User, error) {
	var users []user.User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepo) Update(userUpdate user.User, id string) (user.User, error) {
	var user user.User
	err := r.db.Model(&user).Where("id = ?", id).Updates(&userUpdate).Error
	if err != nil {
		return user, err
	}
	return r.FindByID(id)
}

func (r *userRepo) Delete(id string) (user.User, error) {
	var user user.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepo) FindByEmail(email string) (user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
