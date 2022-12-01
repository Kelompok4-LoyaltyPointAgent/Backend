package transaction_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(query any, args ...any) ([]models.Transaction, error)
	FindByID(id any) (models.Transaction, error)
	Create(transaction models.Transaction) (models.Transaction, error)
	Update(updates models.Transaction, id any) (models.Transaction, error)
	Delete(id any) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) FindAll(query any, args ...any) ([]models.Transaction, error) {
	var transactions []models.Transaction

	var err error
	if len(args) > 0 {
		err = r.db.Where(query, args...).Find(&transactions).Error
	} else {
		err = r.db.Find(&transactions).Error
	}

	return transactions, err
}

func (r *transactionRepository) FindByID(id any) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("id = ?", id).First(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *transactionRepository) Update(updates models.Transaction, id any) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Model(&transaction).Where("id = ?", id).Updates(&updates).Error
	if err != nil {
		return transaction, err
	}
	return r.FindByID(id)
}

func (r *transactionRepository) Delete(id any) error {
	var transaction models.Transaction
	err := r.db.Where("id = ?", id).Delete(&transaction).Error
	return err
}
