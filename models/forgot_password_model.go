package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ForgotPassword struct {
	gorm.Model
	ID        uuid.UUID
	UserID    uuid.UUID
	AccessKey string
	ExpiredAt time.Time
}

func (otp *ForgotPassword) BeforeCreate(tx *gorm.DB) (err error) {
	otp.ID = uuid.New()
	return
}
