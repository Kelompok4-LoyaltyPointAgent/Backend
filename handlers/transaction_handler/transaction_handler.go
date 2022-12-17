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
	GetInvoiceURL(c echo.Context) error
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
	filter := ""

	filterParam := c.QueryParam("type")
	if filterParam == "Purchase" || filterParam == "Redeem" || filterParam == "" {
		filter = filterParam
	} else {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("invalid type"))
	}

	claims := helper.GetTokenClaims(c)

	transactions, err := h.service.FindAllDetail(claims, filter)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, transactions)
}

func (h *transactionHandler) GetTransaction(c echo.Context) error {
	claims := helper.GetTokenClaims(c)

	transaction, err := h.service.FindByID(c.Param("id"), claims)
	if err != nil {
		if err.Error() == "forbidden" {
			return response.Error(c, "failed", http.StatusUnauthorized, err)
		}
		return response.Error(c, "failed", http.StatusNotFound, errors.New("not found"))
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

	transaction, err := h.service.Create(payload, claims)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, transaction)
}

func (h *transactionHandler) UpdateTransaction(c echo.Context) error {
	var payload payload.TransactionUpdatePayload

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
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	_, err := h.service.CallbackXendit(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, nil)
}

func (h *transactionHandler) GetInvoiceURL(c echo.Context) error {
	claims := helper.GetTokenClaims(c)

	url, err := h.service.GetInvoiceURL(c.Param("id"), claims.ID.String())
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, echo.Map{"invoice_url": url})

}
