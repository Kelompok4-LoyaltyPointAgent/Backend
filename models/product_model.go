package models

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID               uuid.UUID
	Name             string
	Description      string
	Type             constant.ProductTypeEnum
	Provider         string
	Price            uint
	PricePoints      uint
	RewardPoints     uint
	Stock            uint
	Recommended      bool
	ProductPictureID *uuid.UUID
	ProductPicture   *ProductPicture
	IconID           *uuid.UUID
	Icon             *ProductPicture
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}
