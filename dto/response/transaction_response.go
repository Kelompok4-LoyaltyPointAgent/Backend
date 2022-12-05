package response

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/models"
)

type TransactionResponse struct {
	ID                uuid.UUID                      `json:"id"`
	UserID            uuid.UUID                      `json:"user_id"`
	User              *UserResponse                  `json:"user,omitempty"`
	ProductID         uuid.UUID                      `json:"product_id"`
	Product           *ProductResponse               `json:"product,omitempty"`
	Amount            float64                        `json:"amount"`
	Method            string                         `json:"method"`
	Status            constant.TransactionStatusEnum `json:"status"`
	Type              constant.TransactionTypeEnum   `json:"type"`
	InvoiceURL        string                         `json:"payout_url,omitempty"`
	TransactionDetail *TransactionDetailResponse     `json:"transaction_detail,omitempty"`
}

type TransactionDetailResponse struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	Number        string    `json:"number"`
	Email         string    `json:"email"`
}

func NewTransactionDetailResponse(transactionDetail models.TransactionDetail) *TransactionDetailResponse {
	return &TransactionDetailResponse{
		ID:            transactionDetail.ID,
		TransactionID: transactionDetail.TransactionID,
		Number:        transactionDetail.Number,
		Email:         transactionDetail.Email,
	}
}

func NewTransactionResponse(transaction models.Transaction, transactionDetail models.TransactionDetail, invoiceURL string) *TransactionResponse {

	response := &TransactionResponse{
		ID:                transaction.ID,
		UserID:            transaction.UserID,
		Amount:            transaction.Amount,
		Method:            transaction.Method,
		Status:            transaction.Status,
		Type:              transaction.Type,
		TransactionDetail: NewTransactionDetailResponse(transactionDetail),
	}

	if invoiceURL != "" {
		response.InvoiceURL = invoiceURL
	}

	if transaction.Product != nil {
		response.ProductID = *transaction.ProductID
		response.Product = NewProductResponse(*transaction.Product)
	}

	if transaction.User != nil {
		response.User = NewUserResponse(*transaction.User)
	}

	return response
}

func NewTransactionsResponse(transactions []models.Transaction) *[]TransactionResponse {
	var response []TransactionResponse
	for _, td := range transactions {
		response = append(response, *NewTransactionResponse(td, *td.TransactionDetail, ""))
	}
	return &response
}
