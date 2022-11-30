package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductPicture struct {
	gorm.Model
	ID   uuid.UUID
	Name string
	Url  string
}

func (productPicture *ProductPicture) BeforeCreate(tx *gorm.DB) (err error) {
	productPicture.ID = uuid.New()
	return
}
