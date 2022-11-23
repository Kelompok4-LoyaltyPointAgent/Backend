package credit_handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/credit_service"
	"github.com/labstack/echo/v4"
)

type CreditHandler interface {
	GetCredits(c echo.Context) error
	CreateCredit(c echo.Context) error
}

type creditHandler struct {
	validate *validator.Validate
	service  credit_service.CreditService
}

func NewCreditHandler(service credit_service.CreditService) CreditHandler {
	validate := validator.New()
	return &creditHandler{validate, service}
}

func (h *creditHandler) GetCredits(c echo.Context) error {
	creditsResponse, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, creditsResponse)
}

func (h *creditHandler) CreateCredit(c echo.Context) error {
	var payload payload.CreditPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	creditResponse, err := h.service.CreateCredit(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, creditResponse)
}
