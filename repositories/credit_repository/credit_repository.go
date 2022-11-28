package credit_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type CreditRepository interface {
	FindAll() ([]models.Credit, error)
	FindByProductID(id any) (models.Credit, error)
	Create(credit models.Credit) (models.Credit, error)
	UpdateByProductID(creditUpdate models.Credit, productID any) (models.Credit, error)
	DeleteByProductID(productID any) error
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

func (r *creditRepository) FindByProductID(id any) (models.Credit, error) {
	var credit models.Credit
	err := r.db.Where("product_id = ?", id).First(&credit).Error
	return credit, err
}

func (r *creditRepository) Create(credit models.Credit) (models.Credit, error) {
	err := r.db.Create(&credit).Error
	return credit, err
}

func (r *creditRepository) UpdateByProductID(creditUpdate models.Credit, productID any) (models.Credit, error) {
	var credit models.Credit
	err := r.db.Model(&credit).Where("product_id = ?", productID).Updates(&creditUpdate).Error
	if err != nil {
		return credit, err
	}
	return r.FindByProductID(productID)
}

func (r *creditRepository) DeleteByProductID(productID any) error {
	var credit models.Credit
	err := r.db.Where("product_id = ?", productID).Delete(&credit).Error
	return err
}
