package forgot_password_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type ForgotPasswordRepository interface {
	FindByID(id any) (models.ForgotPassword, error)
	FindByAccessKey(accessKey any) (models.ForgotPassword, error)
	FindByAccessKeyAndUserID(accessKey any, userID any) (models.ForgotPassword, error)
	Create(forgotPassword models.ForgotPassword) (models.ForgotPassword, error)
	Update(updates models.ForgotPassword, id any) (models.ForgotPassword, error)
	Delete(id any) error
}

type forgotPasswordRepository struct {
	db *gorm.DB
}

func NewForgotPasswordRepository(db *gorm.DB) ForgotPasswordRepository {
	return &forgotPasswordRepository{db}
}

func (r *forgotPasswordRepository) FindByID(id any) (models.ForgotPassword, error) {
	var forgotPassword models.ForgotPassword
	err := r.db.Where("id = ?", id).First(&forgotPassword).Error
	return forgotPassword, err
}

func (r *forgotPasswordRepository) FindByAccessKey(accessKey any) (models.ForgotPassword, error) {
	var forgotPassword models.ForgotPassword
	err := r.db.Where("access_key = ?", accessKey).First(&forgotPassword).Error
	return forgotPassword, err
}

func (r *forgotPasswordRepository) FindByAccessKeyAndUserID(accessKey any, userID any) (models.ForgotPassword, error) {
	var forgotPassword models.ForgotPassword
	err := r.db.Where("access_key = ? AND user_id = ?", accessKey, userID).First(&forgotPassword).Error
	return forgotPassword, err
}

func (r *forgotPasswordRepository) Create(forgotPassword models.ForgotPassword) (models.ForgotPassword, error) {
	err := r.db.Create(&forgotPassword).Error
	return forgotPassword, err
}

func (r *forgotPasswordRepository) Update(updates models.ForgotPassword, id any) (models.ForgotPassword, error) {
	var forgotPassword models.ForgotPassword
	err := r.db.Model(&forgotPassword).Where("id = ?", id).Updates(&updates).Error
	if err != nil {
		return forgotPassword, err
	}
	return r.FindByID(id)
}

func (r *forgotPasswordRepository) Delete(id any) error {
	var forgotPassword models.ForgotPassword
	err := r.db.Where("id = ?", id).Delete(&forgotPassword).Error
	return err
}
