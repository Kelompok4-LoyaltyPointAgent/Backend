package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID               uuid.UUID
	Name             string
	Type             string
	Provider         string
	Price            uint
	PricePoints      uint
	RewardPoints     uint
	Stock            uint
	Recommended      bool
	Description      string
	ProductPictureID *uuid.UUID
	ProductPicture   *ProductPicture
	IconID           *uuid.UUID
	Icon             *ProductPicture
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}
