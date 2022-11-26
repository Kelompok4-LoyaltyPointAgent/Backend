package otp_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type OTPRepository interface {
	FindByID(id any) (models.OTP, error)
	FindByPin(pin any) (models.OTP, error)
	FindByPinAndUserID(pin any, userID any) (models.OTP, error)
	Create(otp models.OTP) (models.OTP, error)
	Update(updates models.OTP, id any) (models.OTP, error)
	DeleteByID(id any) error
}

type otpRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) OTPRepository {
	return &otpRepository{db}
}

func (r *otpRepository) FindByID(id any) (models.OTP, error) {
	var otp models.OTP
	err := r.db.Where("id = ?", id).First(&otp).Error
	return otp, err
}

func (r *otpRepository) FindByPin(pin any) (models.OTP, error) {
	var otp models.OTP
	err := r.db.Where("pin = ?", pin).First(&otp).Error
	return otp, err
}

func (r *otpRepository) FindByPinAndUserID(pin any, userID any) (models.OTP, error) {
	var otp models.OTP
	err := r.db.Where("pin = ? AND user_id = ?", pin, userID).First(&otp).Error
	return otp, err
}

func (r *otpRepository) Create(otp models.OTP) (models.OTP, error) {
	err := r.db.Create(&otp).Error
	return otp, err
}

func (r *otpRepository) Update(updates models.OTP, id any) (models.OTP, error) {
	var otp models.OTP
	err := r.db.Model(&otp).Where("id = ?", id).Updates(&updates).Error
	if err != nil {
		return otp, err
	}
	return r.FindByID(id)
}

func (r *otpRepository) DeleteByID(id any) error {
	var otp models.OTP
	err := r.db.Where("id = ?", id).Delete(&otp).Error
	return err
}
