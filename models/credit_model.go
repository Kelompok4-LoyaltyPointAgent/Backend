package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Credit struct {
	gorm.Model
	ID           uuid.UUID
	ProductID    *uuid.UUID
	Product      Product
	ActivePeriod int
	Amount       int
	Call         int
	SMS          int
}

func (credit *Credit) BeforeCreate(tx *gorm.DB) (err error) {
	credit.ID = uuid.New()
	return
}
