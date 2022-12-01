package transaction_service

import (
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/transaction_repository"
)

type TransactionService interface {
	FindAll(query any, args ...any) (*[]response.TransactionResponse, error)
	FindByID(id any) (*response.TransactionResponse, error)
	Create(payload payload.TransactionPayload) (*response.TransactionResponse, error)
	Update(payload payload.TransactionPayload, id any) (*response.TransactionResponse, error)
	Delete(id any) error
	Cancel(id any) (*response.TransactionResponse, error)
}

type transactionService struct {
	transactionRepository transaction_repository.TransactionRepository
	productRepository     product_repository.ProductRepository
}

func NewTransactionService(
	transactionRepository transaction_repository.TransactionRepository,
	productRepository product_repository.ProductRepository,
) TransactionService {
	return &transactionService{transactionRepository, productRepository}
}

func (s *transactionService) FindAll(query any, args ...any) (*[]response.TransactionResponse, error) {
	transactions, err := s.transactionRepository.FindAll(query, args...)
	if err != nil {
		return nil, err
	}

	return response.NewTransactionsResponse(transactions), nil
}

func (s *transactionService) FindByID(id any) (*response.TransactionResponse, error) {
	transaction, err := s.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return response.NewTransactionResponse(transaction), nil
}

func (s *transactionService) Create(payload payload.TransactionPayload) (*response.TransactionResponse, error) {
	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		return nil, err
	}

	product, err := s.productRepository.FindByID(payload.ProductID)
	if err != nil {
		return nil, err
	}

	var amount float64
	// Blank status indicates transaction made by customer.
	if payload.Status == "" {
		if payload.Type == constant.TransactionTypePurchase {
			var adminFee float64 = 1000
			amount = float64(product.Price) + adminFee
		} else if payload.Type == constant.TransactionTypeRedeem {
			amount = float64(product.PricePoints)
		}
	} else {
		amount = payload.Amount
	}

	transaction, err := s.transactionRepository.Create(models.Transaction{
		UserID:        userID,
		ProductID:     product.ID,
		Amount:        amount,
		PaymentMethod: payload.PaymentMethod,
		PhoneNumber:   payload.PhoneNumber,
		Email:         payload.Email,
		Status:        payload.Status,
		Type:          payload.Type,
	})
	if err != nil {
		return nil, err
	}

	// if transaction.Status == "" {
	// 	// TODO: send bill via payment gateway
	// }

	return response.NewTransactionResponse(transaction), nil
}

func (s *transactionService) Update(payload payload.TransactionPayload, id any) (*response.TransactionResponse, error) {
	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		return nil, err
	}

	product, err := s.productRepository.FindByID(payload.ProductID)
	if err != nil {
		return nil, err
	}

	transaction, err := s.transactionRepository.Update(models.Transaction{
		UserID:      userID,
		ProductID:   product.ID,
		PhoneNumber: payload.PhoneNumber,
		Amount:      payload.Amount,
		Email:       payload.Email,
		Status:      payload.Status,
		Type:        payload.Type,
	}, id)
	if err != nil {
		return nil, err
	}

	return response.NewTransactionResponse(transaction), nil
}

func (s *transactionService) Delete(id any) error {
	return s.transactionRepository.Delete(id)
}

func (s *transactionService) Cancel(id any) (*response.TransactionResponse, error) {
	transaction, err := s.transactionRepository.Update(models.Transaction{
		Status: constant.XenditStatusVoided,
	}, id)
	if err != nil {
		return nil, err
	}

	// TODO: cancel pending transaction via payment gateway

	return response.NewTransactionResponse(transaction), nil
}