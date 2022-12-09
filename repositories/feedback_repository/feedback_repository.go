package feedback_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	FindAll() ([]models.Feedbacks, error)
	FindByID(id any) (models.Feedbacks, error)
	Create(feedback models.Feedbacks) (models.Feedbacks, error)
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db}
}

func (r *feedbackRepository) FindAll() ([]models.Feedbacks, error) {
	var feedbacks []models.Feedbacks
	err := r.db.Preload("User").Find(&feedbacks).Error
	return feedbacks, err
}

func (r *feedbackRepository) FindByID(id any) (models.Feedbacks, error) {
	var feedback models.Feedbacks
	err := r.db.Where("id = ?", id).Preload("User").First(&feedback).Error
	return feedback, err
}

func (r *feedbackRepository) Create(feedback models.Feedbacks) (models.Feedbacks, error) {
	err := r.db.Create(&feedback).Error
	return feedback, err
}
