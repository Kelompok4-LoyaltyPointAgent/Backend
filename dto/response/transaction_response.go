package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type TransactionResponse struct {
	ID            uuid.UUID                             `json:"id"`
	UserID        uuid.UUID                             `json:"user_id"`
	User          *UserResponse                         `json:"user,omitempty"`
	ProductID     uuid.UUID                             `json:"product_id"`
	Product       *ProductResponse                      `json:"product,omitempty"`
	Amount        float64                               `json:"amount"`
	PaymentMethod constant.TransactionPaymentMethodEnum `json:"payment_method"`
	PhoneNumber   string                                `json:"phone_number"`
	Email         string                                `json:"email"`
	Status        constant.XenditStatusEnum             `json:"status"`
	// Status        constant.MidtransStatusEnum  `json:"status"`
	Type constant.TransactionTypeEnum `json:"type"`
}

func NewTransactionResponse(transaction models.Transaction) *TransactionResponse {
	response := &TransactionResponse{
		ID:            transaction.ID,
		UserID:        transaction.UserID,
		ProductID:     transaction.ProductID,
		Amount:        transaction.Amount,
		PaymentMethod: transaction.PaymentMethod,
		PhoneNumber:   transaction.PhoneNumber,
		Email:         transaction.Email,
		Status:        transaction.Status,
		Type:          transaction.Type,
	}

	if transaction.Product != nil {
		response.Product = NewProductResponse(*transaction.Product)
	}

	if transaction.User != nil {
		response.User = NewUserResponse(*transaction.User)
	}

	return response
}

func NewTransactionsResponse(transactions []models.Transaction) *[]TransactionResponse {
	var response []TransactionResponse
	for _, transaction := range transactions {
		response = append(response, *NewTransactionResponse(transaction))
	}
	return &response
}
