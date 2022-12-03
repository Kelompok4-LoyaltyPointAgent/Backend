package models

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID                uuid.UUID
	UserID            uuid.UUID
	User              *User
	ProductID         uuid.UUID
	Product           *Product
	Amount            float64
	Method            string
	Status            constant.TransactionStatusEnum
	Type              constant.TransactionTypeEnum
	TransactionDetail *TransactionDetail
}

type TransactionDetail struct {
	gorm.Model
	ID            uuid.UUID
	TransactionID uuid.UUID
	Number        string
	Email         string
}

func (transactionDetail *TransactionDetail) BeforeCreate(tx *gorm.DB) (err error) {
	transactionDetail.ID = uuid.New()
	return
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	transaction.ID = uuid.New()
	return
}
