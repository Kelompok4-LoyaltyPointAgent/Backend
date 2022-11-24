package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type CreditResponse struct {
	ID           uuid.UUID        `json:"id"`
	ProductID    *uuid.UUID       `json:"product_id,omitempty"`
	Product      *ProductResponse `json:"product,omitempty"`
	Description  string           `json:"description"`
	ActivePeriod int              `json:"active_period"`
	Amount       int              `json:"amount"`
}

func NewCreditResponse(credit models.Credit) *CreditResponse {
	response := &CreditResponse{
		ID:           credit.ID,
		ProductID:    credit.ProductID,
		Description:  credit.Description,
		ActivePeriod: credit.ActivePeriod,
		Amount:       credit.Amount,
	}
	if credit.Product != nil {
		response.Product = NewProductResponse(*credit.Product)
	}
	return response
}

func NewCreditsResponse(credits []models.Credit) *[]CreditResponse {
	var response []CreditResponse
	for _, credit := range credits {
		response = append(response, *NewCreditResponse(credit))
	}
	return &response
}
