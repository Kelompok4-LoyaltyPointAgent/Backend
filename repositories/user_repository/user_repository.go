package user_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
	Update(user models.User, id string) (models.User, error)
	Delete(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) Update(userUpdate models.User, id string) (models.User, error) {
	var user models.User
	err := r.db.Model(&user).Where("id = ?", id).Updates(userUpdate).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Delete(id string) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
