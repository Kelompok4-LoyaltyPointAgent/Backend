package cached_payouturl_repository

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type PayoutURLRepository interface {
	GetPayoutURL(transactionID string) (string, error)
	SetPayoutURL(url string, transactionID string) error
	DeletePayoutURL(transactionID string) error
}

type payoutURLRepository struct {
	db *redis.Client
}

func NewPayoutURLRepository(db *redis.Client) PayoutURLRepository {
	return &payoutURLRepository{db: db}
}

func (r *payoutURLRepository) GetPayoutURL(transactionID string) (string, error) {
	key := fmt.Sprintf("transactions:payouturl:%s", transactionID)
	var url string
	if err := r.db.Get(r.db.Context(), key).Scan(&url); err != nil {
		return "", err
	}
	return url, nil
}

func (r *payoutURLRepository) SetPayoutURL(url string, transactionID string) error {
	key := fmt.Sprintf("transactions:payouturl:%s", transactionID)
	return r.db.Set(r.db.Context(), key, url, 0).Err()
}

func (r *payoutURLRepository) DeletePayoutURL(transactionID string) error {
	key := fmt.Sprintf("transactions:payouturl:%s", transactionID)
	return r.db.Del(r.db.Context(), key).Err()
}
