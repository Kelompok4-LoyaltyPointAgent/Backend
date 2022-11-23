package credit_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type CreditRepository interface {
	FindAll() ([]models.Credit, error)
	Create(credit models.Credit) (models.Credit, error)
}

type creditRepository struct {
	db *gorm.DB
}

func NewCreditRepository(db *gorm.DB) CreditRepository {
	return &creditRepository{db}
}

func (r *creditRepository) FindAll() ([]models.Credit, error) {
	var credits []models.Credit
	err := r.db.Preload("Product").Find(&credits).Error
	return credits, err
}

func (r *creditRepository) Create(credit models.Credit) (models.Credit, error) {
	err := r.db.Create(&credit).Error
	return credit, err
}
