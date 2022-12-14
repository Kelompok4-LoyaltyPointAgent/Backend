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
	ProductID         *uuid.UUID                     `json:"product_id"`
	Product           *ProductResponse               `json:"product,omitempty"`
	Amount            float64                        `json:"amount"`
	Method            string                         `json:"method"`
	Status            constant.TransactionStatusEnum `json:"status"`
	Type              constant.TransactionTypeEnum   `json:"type"`
	InvoiceURL        string                         `json:"invoice_url,omitempty"`
	TransactionDetail *TransactionDetailResponse     `json:"transaction_detail,omitempty"`
	CreatedDate       string                         `json:"created_date,omitempty"`
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
		ProductID:         transaction.ProductID,
		TransactionDetail: NewTransactionDetailResponse(transactionDetail),
		CreatedDate:       transaction.CreatedAt.Format("02 January 2006"),
	}

	if invoiceURL != "" {
		response.InvoiceURL = invoiceURL
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
	for _, td := range transactions {
		if td.TransactionDetail != nil {
			response = append(response, *NewTransactionResponse(td, *td.TransactionDetail, ""))
		} else {
			response = append(response, *NewTransactionResponse(td, models.TransactionDetail{}, ""))
		}
	}
	return &response
}
