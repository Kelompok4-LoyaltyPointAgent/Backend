package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	ID        uuid.UUID
	UserID    uuid.UUID
	Pin       string
	ExpiredAt time.Time
}

func (OTP) TableName() string {
	return "one_time_passwords"
}

func (otp *OTP) BeforeCreate(tx *gorm.DB) (err error) {
	otp.ID = uuid.New()
	return
}
