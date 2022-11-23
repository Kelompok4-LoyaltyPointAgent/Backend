package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type CreditResponse struct {
	ID           uuid.UUID        `json:"id"`
	ProductID    *uuid.UUID       `json:"product_id"`
	Product      *ProductResponse `json:"product,omitempty"`
	Description  string           `json:"description"`
	ActivePeriod int              `json:"active_period"`
	Amount       int              `json:"amount"`
}

func NewCreditResponse(credit models.Credit) *CreditResponse {
	return &CreditResponse{
		ID:           credit.ID,
		ProductID:    credit.ProductID,
		Product:      NewProductResponse(*credit.Product),
		Description:  credit.Description,
		ActivePeriod: credit.ActivePeriod,
		Amount:       credit.Amount,
	}
}

func NewCreditsResponse(credits []models.Credit) *[]CreditResponse {
	var response []CreditResponse
	for _, credit := range credits {
		response = append(response, *NewCreditResponse(credit))
	}
	return &response
}
