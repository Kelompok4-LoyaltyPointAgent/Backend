package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Packages struct {
	gorm.Model
	ID           uuid.UUID
	ProductID    *uuid.UUID
	Product      *Product
	ActivePeriod int
	Internet     int
	Call         int
}

func (packages *Packages) BeforeCreate(tx *gorm.DB) (err error) {
	packages.ID = uuid.New()
	return
}
