package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FAQ struct {
	gorm.Model
	ID       uuid.UUID
	Category string
	Question string
	Answer   string
}

func (FAQ) TableName() string {
	return "frequently_asked_questions"
}

func (faq *FAQ) BeforeCreate(tx *gorm.DB) (err error) {
	faq.ID = uuid.New()
	return
}
