package cached_analytics_repository

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/analytics_repository"
)

type CachedAnalyticsRepository interface {
	analytics_repository.AnalyticsRepository
	CheckSalesCount(year int) bool
	SetSalesCount(year, count int) error
	CheckIncome(year int) bool
	SetIncome(year int, sum float64) error
	CheckTransactionsByMonth(year int) bool
	SetTransactionsByMonth(year int, data string) error
	CheckTransactionsByType(year int) bool
	SetTransactionsByType(year int, data string) error
	CheckRecentTransactions(year int) bool
	SetRecentTransactions(year int, data string) error
	GetDataInStock() (*string, error)
	CheckDataInStock() bool
	SetProductCount(data string) error
}

type cachedAnalyticsRepository struct {
	db *redis.Client
}

func NewCachedAnalyticsRepository(db *redis.Client) CachedAnalyticsRepository {
	return &cachedAnalyticsRepository{db}
}

func (r *cachedAnalyticsRepository) GetDataInStock() (*string, error) {
	key := "analytics:stock_data"
	var data string
	if err := r.db.Get(r.db.Context(), key).Scan(&data); err != nil {
		return nil, err
	}
	log.Println("------------")
	log.Println(data)
	return &data, nil
}

func (r *cachedAnalyticsRepository) SetProductCount(data string) error {
	key := "analytics:stock_data"
	exp := time.Duration(30 * time.Second)
	return r.db.Set(r.db.Context(), key, data, exp).Err()
}

func (r *cachedAnalyticsRepository) CheckDataInStock() bool {
	key := "analytics:stock_data"
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}

func (r *cachedAnalyticsRepository) ProductCount() (int, error) {
	return 0, nil
}

func (r *cachedAnalyticsRepository) SalesCount(year int) int {
	key := fmt.Sprintf("analytics:sales_count:%d", year)
	count := 0
	if err := r.db.Get(r.db.Context(), key).Scan(&count); err != nil {
		return 0
	}
	return count
}

func (r *cachedAnalyticsRepository) CheckSalesCount(year int) bool {
	key := fmt.Sprintf("analytics:sales_count:%d", year)
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}

func (r *cachedAnalyticsRepository) SetSalesCount(year, count int) error {
	key := fmt.Sprintf("analytics:sales_count:%d", year)
	exp := time.Duration(30 * time.Second)
	return r.db.Set(r.db.Context(), key, count, exp).Err()
}

func (r *cachedAnalyticsRepository) Income(year int) float64 {
	key := fmt.Sprintf("analytics:income:%d", year)
	var sum float64
	if err := r.db.Get(r.db.Context(), key).Scan(&sum); err != nil {
		return 0
	}
	return sum
}

func (r *cachedAnalyticsRepository) CheckIncome(year int) bool {
	key := fmt.Sprintf("analytics:income:%d", year)
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}

func (r *cachedAnalyticsRepository) SetIncome(year int, sum float64) error {
	key := fmt.Sprintf("analytics:income:%d", year)
	exp := time.Duration(30 * time.Second)
	return r.db.Set(r.db.Context(), key, sum, exp).Err()
}

func (r *cachedAnalyticsRepository) TransactionsByMonth(year int) analytics_repository.TransactionsByMonth {
	key := fmt.Sprintf("analytics:transactions_by_month:%d", year)
	var transactions analytics_repository.TransactionsByMonth
	var data string
	if err := r.db.Get(r.db.Context(), key).Scan(&data); err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(data), &transactions); err != nil {
		log.Printf("JSON error: %s", err)
		return nil
	}
	return transactions
}

func (r *cachedAnalyticsRepository) CheckTransactionsByMonth(year int) bool {
	key := fmt.Sprintf("analytics:transactions_by_month:%d", year)
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}

func (r *cachedAnalyticsRepository) SetTransactionsByMonth(year int, data string) error {
	key := fmt.Sprintf("analytics:transactions_by_month:%d", year)
	exp := time.Duration(30 * time.Second)
	return r.db.Set(r.db.Context(), key, data, exp).Err()
}

func (r *cachedAnalyticsRepository) TransactionsByType(year int) analytics_repository.TransactionsByType {
	key := fmt.Sprintf("analytics:transactions_by_type:%d", year)
	var transactions analytics_repository.TransactionsByType
	var data string
	if err := r.db.Get(r.db.Context(), key).Scan(&data); err != nil {
		return nil
	}
	if err := json.Unmarshal([]byte(data), &transactions); err != nil {
		log.Printf("JSON error: %s", err)
		return nil
	}
	return transactions
}

func (r *cachedAnalyticsRepository) CheckTransactionsByType(year int) bool {
	key := fmt.Sprintf("analytics:transactions_by_type:%d", year)
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}

func (r *cachedAnalyticsRepository) SetTransactionsByType(year int, data string) error {
	key := fmt.Sprintf("analytics:transactions_by_type:%d", year)
	exp := time.Duration(30 * time.Second)
	return r.db.Set(r.db.Context(), key, data, exp).Err()
}

func (r *cachedAnalyticsRepository) RecentTransactions(year, limit int) ([]models.Transaction, error) {
	key := fmt.Sprintf("analytics:recent_transactions:%d", year)
	var transactions []models.Transaction
	var data string
	if err := r.db.Get(r.db.Context(), key).Scan(&data); err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &transactions); err != nil {
		log.Printf("JSON error: %s", err)
		return nil, err
	}
	return transactions[:limit], nil
}

func (r *cachedAnalyticsRepository) CheckRecentTransactions(year int) bool {
	key := fmt.Sprintf("analytics:recent_transactions:%d", year)
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}

func (r *cachedAnalyticsRepository) SetRecentTransactions(year int, data string) error {
	key := fmt.Sprintf("analytics:recent_transactions:%d", year)
	exp := time.Duration(30 * time.Second)
	return r.db.Set(r.db.Context(), key, data, exp).Err()
}
