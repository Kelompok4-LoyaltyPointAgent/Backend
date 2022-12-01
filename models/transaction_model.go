package models

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            uuid.UUID
	UserID        uuid.UUID
	User          *User
	ProductID     uuid.UUID
	Product       *Product
	Amount        float64
	PaymentMethod constant.TransactionPaymentMethodEnum
	PhoneNumber   string
	Email         string
	Status        constant.XenditStatusEnum
	Type          constant.TransactionTypeEnum
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	transaction.ID = uuid.New()
	return
}
