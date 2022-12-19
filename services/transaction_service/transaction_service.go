package transaction_service

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/cachedrepositories/cached_invoiceurl_repository"
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
	Update(payload payload.TransactionUpdatePayload, id any) (*response.TransactionResponse, error)
	Delete(id any) error
	Cancel(id any) (*response.TransactionResponse, error)
	CallbackXendit(payload map[string]interface{}) (bool, error)
	GetInvoiceURL(id string, userId string) (*string, error)
}

type transactionService struct {
	transactionRepository      transaction_repository.TransactionRepository
	productRepository          product_repository.ProductRepository
	userRepository             user_repository.UserRepository
	cachedInvoiceURLRepository cached_invoiceurl_repository.InvoiceURLRepository
}

func NewTransactionService(
	transactionRepository transaction_repository.TransactionRepository,
	productRepository product_repository.ProductRepository,
	userRepository user_repository.UserRepository,
	cachedInvoiceURL cached_invoiceurl_repository.InvoiceURLRepository,
) TransactionService {
	return &transactionService{transactionRepository, productRepository, userRepository, cachedInvoiceURL}
}

func (s *transactionService) FindAllDetail(claims *helper.JWTCustomClaims, filter any) (*[]response.TransactionResponse, error) {

	var transactions []models.Transaction
	var err error
	if claims.Role == constant.UserRoleAdmin.String() {
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
	if claims.Role == constant.UserRoleAdmin.String() {
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
	var product models.Product

	if claims.Role != constant.UserRoleAdmin.String() {
		payload.UserID = claims.ID.String()
		payload.Status = ""
	}

	if payload.ProductID != "" {
		getProduct, err := s.productRepository.FindByID(payload.ProductID)
		if err != nil {
			return nil, err
		}
		product = getProduct
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
			if product.Stock < 1 {
				return nil, errors.New("product has not enough stocks")
			}

			payload.Method = ""
			var adminFee float64 = 1000
			amount = float64(product.Price) + adminFee
			payload.Status = constant.TransactionStatusPending

		} else if payload.Type == constant.TransactionTypeRedeem {

			if product.Stock < 1 {
				return nil, errors.New("product has not enough stocks")
			}

			payload.Method = ""
			amount = float64(product.PricePoints)

			if user.Points < product.PricePoints {
				return nil, errors.New("user has not enough points")
			}

			points := user.Points - product.PricePoints

			if _, err := s.userRepository.UpdateUserPoint(points, user.ID.String()); err != nil {
				return nil, err
			}

			stock := product.Stock - 1

			if _, err := s.productRepository.UpdateStockProduct(stock, product.ID); err != nil {
				return nil, err
			}

			payload.Status = constant.TransactionStatusSuccess
		} else if payload.Type == constant.TransactionTypeCashout {
			if user.Points < 50000 {
				return nil, errors.New("user has not enough points")
			}
			payload.Status = constant.TransactionStatusPending
			payload.ProductID = ""
			amount = payload.Amount

		}
	} else {
		amount = payload.Amount
	}

	transactionModel := models.Transaction{
		UserID: user.ID,
		Amount: amount,
		Type:   payload.Type,
		Method: payload.Method,
		Status: payload.Status,
		TransactionDetail: &models.TransactionDetail{
			Email:  payload.Email,
			Number: payload.Number,
		},
	}

	if product.ID != uuid.Nil {
		transactionModel.ProductID = &product.ID
	} else {
		transactionModel.ProductID = nil
	}

	transaction, err := s.transactionRepository.Create(transactionModel)

	if err != nil {
		return nil, err
	}

	if transaction.Status == constant.TransactionStatusPending && claims.Role != constant.UserRoleAdmin.String() && transaction.Type == constant.TransactionTypePurchase {
		resp, err := helper.CreateInvoiceXendit(transaction, *transaction.TransactionDetail, user)
		if err != nil {
			return nil, err
		}

		//Save invoice url in redis
		if err := s.cachedInvoiceURLRepository.SetInvoiceURL(resp.InvoiceURL, transaction.ID.String()); err != nil {
			return nil, err
		}

		return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, resp.InvoiceURL), nil

	} else if transaction.Status == constant.TransactionStatusPending && claims.Role != constant.UserRoleAdmin.String() && transaction.Type == constant.TransactionTypeCashout {

		_, err := helper.CreateDisbursementXendit(transaction, *transaction.TransactionDetail, user)
		if err != nil {
			return nil, err
		}

		return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, ""), nil
	}

	return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, ""), nil
}

func (s *transactionService) Update(payload payload.TransactionUpdatePayload, id any) (*response.TransactionResponse, error) {

	transaction, err := s.transactionRepository.Update(models.Transaction{
		Amount: payload.Amount,
		Status: payload.Status,
		Type:   payload.Type,
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
	err := s.transactionRepository.Delete(id)
	if err != nil {
		return err
	}

	err = s.transactionRepository.DeleteDetailByTransactionID(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *transactionService) Cancel(id any) (*response.TransactionResponse, error) {
	transaction, err := s.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if transaction.Status == constant.TransactionStatusSuccess {
		return nil, errors.New("transaction already success")
	}

	transaction, err = s.transactionRepository.Update(models.Transaction{
		Status: constant.TransactionStatusFailed,
	}, id)
	if err != nil {
		return nil, err
	}

	return response.NewTransactionResponse(transaction, *transaction.TransactionDetail, ""), nil
}

func (s *transactionService) GetInvoiceURL(id string, userId string) (*string, error) {
	//Get Transaction ID
	transaction, err := s.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if transaction.UserID.String() != userId {
		return nil, errors.New("forbidden")
	}

	url, err := s.cachedInvoiceURLRepository.GetInvoiceURL(transaction.ID.String())
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (s *transactionService) CallbackXendit(payload map[string]interface{}) (bool, error) {
	transaction, err := s.transactionRepository.FindByID(payload["external_id"])
	if err != nil {
		return false, err
	}

	if transaction.Status == constant.TransactionStatusSuccess {
		return false, nil
	}

	// Find User ID
	user, err := s.userRepository.FindByID(transaction.UserID.String())
	if err != nil {
		return false, err
	}

	if _, ok := payload["disbursement_description"]; ok {

		if payload["status"] == constant.XenditStatusCompleted.String() {
			transaction.Status = constant.TransactionStatusSuccess

			// Update User Points
			points := user.Points - uint(transaction.Amount)

			if _, err := s.userRepository.UpdateUserPoint(points, user.ID.String()); err != nil {
				return false, err
			}

		} else if payload["status"] == constant.XenditStatusFailed.String() {
			transaction.Status = constant.TransactionStatusFailed

		} else if payload["status"] == constant.XenditStatusPending.String() {
			transaction.Status = constant.TransactionStatusPending
		}
	} else {
		//Callback Xendit Invoice
		if payload["status"].(string) == constant.XenditStatusPaid.String() {
			transaction.Status = constant.TransactionStatusSuccess
			transaction.Method = payload["payment_channel"].(string)

			log.Println("MASUKK")

			// Find Product ID
			product, err := s.productRepository.FindByID(transaction.ProductID.String())
			if err != nil {
				return false, err
			}

			stock := product.Stock - 1

			if _, err := s.productRepository.UpdateStockProduct(stock, product.ID); err != nil {
				return false, err
			}

			// Update User Points
			points := user.Points - product.PricePoints

			if _, err := s.userRepository.UpdateUserPoint(points, user.ID.String()); err != nil {
				return false, err
			}

			//Delete invoice url
			if err := s.cachedInvoiceURLRepository.DeleteInvoiceURL(transaction.ID.String()); err != nil {
				return false, err
			}

		} else if payload["status"] == constant.XenditStatusExpired {
			transaction.Status = constant.TransactionStatusFailed
		}
	}

	if _, err := s.transactionRepository.Update(transaction, transaction.ID.String()); err != nil {
		return false, err
	}

	return true, nil
}
