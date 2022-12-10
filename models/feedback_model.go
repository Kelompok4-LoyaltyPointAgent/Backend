package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedbacks struct {
	gorm.Model
	ID                   uuid.UUID
	UserID               uuid.UUID
	User                 *User
	IsInformationHelpful *bool
	IsArticleHelpful     *bool
	IsArticleEasyToFind  *bool
	IsDesignGood         *bool
	Review               string
}

func (Feedbacks) TableName() string {
	return "feedbacks"
}

func (feedback *Feedbacks) BeforeCreate(tx *gorm.DB) (err error) {
	feedback.ID = uuid.New()
	return
}
