package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Packages struct {
	gorm.Model
	ID            uuid.UUID
	ProductID     *uuid.UUID
	Product       Product
	ActivePeriod  int
	TotalInternet float64
	MainInternet  float64
	NightInternet float64
	SocialMedia   float64
	Call          int
	SMS           int
	Description   string
}

func (packages *Packages) BeforeCreate(tx *gorm.DB) (err error) {
	packages.ID = uuid.New()
	return
}
