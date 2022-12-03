package transaction_service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/constant"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/product_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/transaction_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
)

type TransactionService interface {
	FindAllDetail(claims *helper.JWTCustomClaims, filter any) (*[]response.TransactionResponse, error)
	FindByID(id any, claims *helper.JWTCustomClaims) (*response.TransactionResponse, error)
	Create(payload payload.TransactionPayload, claims *helper.JWTCustomClaims) (*response.TransactionResponse, error)
	Update(payload payload.TransactionPayload, id any) (*response.TransactionResponse, error)
	Delete(id any) error
	Cancel(id any) (*response.TransactionResponse, error)
	CallbackXendit(payload map[string]interface{}) (bool, error)
}

type transactionService struct {
	transactionRepository transaction_repository.TransactionRepository
	productRepository     product_repository.ProductRepository
	userRepository        user_repository.UserRepository
}

func NewTransactionService(
	transactionRepository transaction_repository.TransactionRepository,
	productRepository product_repository.ProductRepository,
	userRepository user_repository.UserRepository,
) TransactionService {
	return &transactionService{transactionRepository, productRepository, userRepository}
}

func (s *transactionService) FindAllDetail(claims *helper.JWTCustomClaims, filter any) (*[]response.TransactionResponse, error) {

	var transactions []models.Transaction
	var err error
	if claims.Role == "Admin" {
		transactions, err = s.transactionRepository.FindAll("", "")
		if err != nil {
			return nil, err
		}
	} else {
		var args []any
		args = append(args, claims.ID.String())
		var query string
		if filter == constant.TransactionTypePurchase.String() {
			args = append(args, filter)
			query = "user_id = ? AND type = ?"
		} else if filter == constant.TransactionTypeRedeem.String() {
			args = append(args, filter)
			query = "user_id = ? AND type = ? OR type = 'Cashout'"
		} else {
			query = "user_id = ?"
		}

		transactions, err = s.transactionRepository.FindAll(query, args...)
		if err != nil {
			return nil, err
		}
	}

	return response.NewTransactionsResponse(transactions), err
}

func (s *transactionService) FindByID(id any, claims *helper.JWTCustomClaims) (*response.TransactionResponse, error) {
	var transaction models.Transaction
	var err error
	if claims.Role == "Admin" {
		transaction, err = s.transactionRepository.FindByID(id)
		if err != nil {
			return nil, err
		}
	} else {
		transaction, err = s.transactionRepository.FindByID(id)
		if err != nil {
			return nil, err
		}
		if transaction.UserID.String() != claims.ID.String() {
			return nil, errors.New("forbidden")
		}
	}

	return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, ""), nil
}

func (s *transactionService) Create(payload payload.TransactionPayload, claims *helper.JWTCustomClaims) (*response.TransactionResponse, error) {
	if claims.Role != "Admin" {
		payload.UserID = claims.ID.String()
		payload.Status = ""
	}

	product, err := s.productRepository.FindByID(payload.ProductID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByID(payload.UserID)
	if err != nil {
		return nil, err
	}

	var amount float64
	// Blank status indicates transaction made by customer.
	if payload.Status == "" {
		if payload.Email == "" {
			payload.Email = user.Email
		}
		if payload.Type == constant.TransactionTypePurchase {
			var adminFee float64 = 1000
			amount = float64(product.Price) + adminFee
			payload.Status = constant.TransactionStatusPending

		} else if payload.Type == constant.TransactionTypeRedeem {
			amount = float64(product.PricePoints)

			if user.Points < product.PricePoints {
				return nil, errors.New("user has not enough points")
			}

			updates := models.User{
				Points: user.Points - product.PricePoints,
			}
			// TODO: make sure user points can be updated to 0
			if _, err := s.userRepository.Update(updates, user.ID.String()); err != nil {
				return nil, err
			}

			payload.Status = constant.TransactionStatusSuccess
		}
	} else {
		amount = payload.Amount
	}

	transaction, err := s.transactionRepository.Create(models.Transaction{
		UserID:    user.ID,
		ProductID: product.ID,
		Amount:    amount,
		Status:    payload.Status,
		Type:      payload.Type,
	})

	transactionDetail, err := s.transactionRepository.CreateDetail(models.TransactionDetail{
		TransactionID: transaction.ID,
		Email:         payload.Email,
		Number:        payload.Number,
	})

	if err != nil {
		return nil, err
	}

	if transaction.Status == constant.TransactionStatusPending && claims.Role != "Admin" {
		// TODO: send bill via payment gateway
		resp, err := helper.CreateInvoiceXendit(transaction, transactionDetail, user)
		if err != nil {
			return nil, err
		}

		return response.NewTransactionResponse(transaction, transactionDetail, resp.InvoiceURL), nil
	}

	return response.NewTransactionResponse(transaction, transactionDetail, ""), nil
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
		UserID:    userID,
		ProductID: product.ID,
		Amount:    payload.Amount,
		Status:    payload.Status,
		Type:      payload.Type,
	}, id)
	if err != nil {
		return nil, err
	}

	transaction, err = s.transactionRepository.FindByID(transaction.ID.String())
	if err != nil {
		return nil, err
	}

	return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, ""), nil
}

func (s *transactionService) Delete(id any) error {
	return s.transactionRepository.Delete(id)
}

func (s *transactionService) Cancel(id any) (*response.TransactionResponse, error) {
	transaction, err := s.transactionRepository.Update(models.Transaction{
		Status: constant.TransactionStatusFailed,
	}, id)
	if err != nil {
		return nil, err
	}

	transaction, err = s.transactionRepository.FindByID(transaction.ID.String())
	if err != nil {
		return nil, err
	}

	// TODO: cancel pending transaction via payment gateway

	return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, ""), nil
}

func (s *transactionService) CallbackXendit(payload map[string]interface{}) (bool, error) {
	transaction, err := s.transactionRepository.FindByID(payload["external_id"])
	if err != nil {
		return false, err
	}

	if transaction.Status == constant.TransactionStatusSuccess {
		return false, nil
	}

	if payload["status"].(string) == constant.XenditStatusPaid.String() {
		transaction.Status = constant.TransactionStatusSuccess
		transaction.Method = payload["payment_channel"].(string)

		// Find User ID
		user, err := s.userRepository.FindByID(transaction.UserID.String())
		if err != nil {
			return false, err
		}

		// Find Product ID
		product, err := s.productRepository.FindByID(transaction.ProductID.String())
		if err != nil {
			return false, err
		}

		// Update User Points
		updates := models.User{
			Points: user.Points + product.PricePoints,
		}

		if _, err := s.userRepository.Update(updates, user.ID.String()); err != nil {
			return false, err
		}

	} else if payload["status"] == constant.XenditStatusExpired {
		transaction.Status = constant.TransactionStatusFailed
	}

	if _, err := s.transactionRepository.Update(transaction, transaction.ID.String()); err != nil {
		return false, err
	}

	return true, nil
}
