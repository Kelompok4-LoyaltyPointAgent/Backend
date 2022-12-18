package analytics_repository

import (
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"gorm.io/gorm"
)

type TransactionsByMonth []struct {
	Month int
	Value int
}

type TransactionsByType []struct {
	Type  string
	Value int
}

type AnalyticsRepository interface {
	SalesCount(year int) int
	Income(year int) float64
	TransactionsByMonth(year int) TransactionsByMonth
	TransactionsByType(year int) TransactionsByType
	RecentTransactions(year, limit int) ([]models.Transaction, error)
	ProductCount() (int, error)
}

type analyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) AnalyticsRepository {
	return &analyticsRepository{db}
}

func (r *analyticsRepository) ProductCount() (int, error) {
	query := "SELECT COUNT(name) FROM products"
	count := 0
	err := r.db.Raw(query).Scan(&count).Error
	return count, err
}

func (r *analyticsRepository) SalesCount(year int) int {
	query := "SELECT COUNT(1) FROM transactions WHERE status=? AND type=? AND YEAR(created_at)=?"
	count := 0
	r.db.Raw(query, constant.TransactionStatusSuccess, constant.TransactionTypePurchase, year).Scan(&count)
	return count
}

func (r *analyticsRepository) Income(year int) float64 {
	query := "SELECT SUM(amount) FROM transactions WHERE status=? AND type=? AND YEAR(created_at)=?"
	var sum float64
	r.db.Raw(query, constant.TransactionStatusSuccess, constant.TransactionTypePurchase, year).Scan(&sum)
	return sum
}

func (r *analyticsRepository) TransactionsByMonth(year int) TransactionsByMonth {
	query := "SELECT MONTH(created_at) AS month, COUNT(1) AS value FROM transactions WHERE status=? AND YEAR(created_at)=? GROUP BY MONTH(created_at) ORDER BY MONTH(created_at) ASC"
	var transactions TransactionsByMonth
	r.db.Raw(query, constant.TransactionStatusSuccess, year).Scan(&transactions)
	return transactions
}

func (r *analyticsRepository) TransactionsByType(year int) TransactionsByType {
	query := "SELECT type, COUNT(1) AS value FROM transactions WHERE status=? AND YEAR(created_at)=? GROUP BY type"
	var transactions TransactionsByType
	r.db.Raw(query, constant.TransactionStatusSuccess, year).Scan(&transactions)
	return transactions
}

func (r *analyticsRepository) RecentTransactions(year, limit int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("YEAR(created_at) = ?", year).Limit(limit).Order("created_at DESC").Preload("TransactionDetail").Preload("Product").Find(&transactions).Error
	return transactions, err
}
