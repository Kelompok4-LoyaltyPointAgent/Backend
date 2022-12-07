package models

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"gorm.io/gorm"
)

type ProductPicture struct {
	gorm.Model
	ID   uuid.UUID
	Name string
	Url  string
	Type constant.ProductPictureTypeEnum
}

func (productPicture *ProductPicture) BeforeCreate(tx *gorm.DB) (err error) {
	productPicture.ID = uuid.New()
	return
}
