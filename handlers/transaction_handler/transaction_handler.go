package transaction_handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/services/transaction_service"
	"github.com/labstack/echo/v4"
)

type TransactionHandler interface {
	GetTransactions(c echo.Context) error
	GetTransaction(c echo.Context) error
	CreateTransaction(c echo.Context) error
	UpdateTransaction(c echo.Context) error
	DeleteTransaction(c echo.Context) error
	CancelTransaction(c echo.Context) error
	TransactionWebhook(c echo.Context) error
}

type transactionHandler struct {
	validate *validator.Validate
	service  transaction_service.TransactionService
}

func NewTransactionHandler(service transaction_service.TransactionService) TransactionHandler {
	validate := validator.New()
	return &transactionHandler{validate, service}
}

func (h *transactionHandler) GetTransactions(c echo.Context) error {
	claims := helper.GetTokenClaims(c)

	var query string
	var args []any
	if claims.Role != "Admin" {
		query = "user_id = ?"
		args = append(args, claims.ID)
	}

	transactions, err := h.service.FindAll(query, args...)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, transactions)
}

func (h *transactionHandler) GetTransaction(c echo.Context) error {
	transaction, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusNotFound, errors.New("not found"))
	}

	claims := helper.GetTokenClaims(c)
	if transaction.UserID != claims.ID {
		return response.Error(c, "failed", http.StatusForbidden, errors.New("forbidden"))
	}

	return response.Success(c, "success", http.StatusOK, transaction)
}

func (h *transactionHandler) CreateTransaction(c echo.Context) error {
	var payload payload.TransactionPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	claims := helper.GetTokenClaims(c)
	if claims.Role != "Admin" {
		payload.UserID = claims.ID.String()
		payload.Status = ""
	}

	transaction, err := h.service.Create(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, transaction)
}

func (h *transactionHandler) UpdateTransaction(c echo.Context) error {
	var payload payload.TransactionPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	transaction, err := h.service.Update(payload, c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, transaction)
}

func (h *transactionHandler) DeleteTransaction(c echo.Context) error {
	if err := h.service.Delete(c.Param("id")); err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, nil)
}

func (h *transactionHandler) CancelTransaction(c echo.Context) error {
	transaction, err := h.service.Cancel(c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, transaction)
}

func (h *transactionHandler) TransactionWebhook(c echo.Context) error {
	// TODO: handle payment gateway request

	return response.Success(c, "success", http.StatusOK, nil)
}
