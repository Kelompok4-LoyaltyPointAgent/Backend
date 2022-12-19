package cached_invoiceurl_repository

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type InvoiceURLRepository interface {
	GetInvoiceURL(transactionID string) (string, error)
	SetInvoiceURL(url string, transactionID string) error
	DeleteInvoiceURL(transactionID string) error
	CheckInvoiceURL(transactionID string) bool
}

type invoiceURLRepository struct {
	db *redis.Client
}

func NewInvoiceURLRepository(db *redis.Client) InvoiceURLRepository {
	return &invoiceURLRepository{db: db}
}

func (r *invoiceURLRepository) GetInvoiceURL(transactionID string) (string, error) {
	key := fmt.Sprintf("transactions:invoiceurl:%s", transactionID)
	var url string
	if err := r.db.Get(r.db.Context(), key).Scan(&url); err != nil {
		return "", err
	}
	return url, nil
}

func (r invoiceURLRepository) SetInvoiceURL(url string, transactionID string) error {
	key := fmt.Sprintf("transactions:invoiceurl:%s", transactionID)
	exp := time.Duration(1 * time.Hour)
	return r.db.Set(r.db.Context(), key, url, exp).Err()
}

func (r *invoiceURLRepository) DeleteInvoiceURL(transactionID string) error {
	key := fmt.Sprintf("transactions:invoiceurl:%s", transactionID)
	return r.db.Del(r.db.Context(), key).Err()
}

func (r *invoiceURLRepository) CheckInvoiceURL(transactionID string) bool {
	key := fmt.Sprintf("transactions:invoiceurl:%s", transactionID)
	result, err := r.db.Exists(r.db.Context(), key).Result()
	if err != nil {
		log.Printf("Redis error: %s", err)
		return false
	}
	return result > 0
}
