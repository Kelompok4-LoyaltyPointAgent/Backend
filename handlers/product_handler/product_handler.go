package product_handler

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/product_service"
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	GetProductsWithCredits(c echo.Context) error
	GetProductWithCredit(c echo.Context) error
	CreateProductWithCredit(c echo.Context) error
	UpdateProductWithCredit(c echo.Context) error
	DeleteProductWithCredit(c echo.Context) error
}

type productHandler struct {
	validate *validator.Validate
	service  product_service.ProductService
}

func NewProductHandler(service product_service.ProductService) ProductHandler {
	validate := validator.New()
	return &productHandler{validate, service}
}

func (h *productHandler) GetProductsWithCredits(c echo.Context) error {
	productsWithCreditsResponse, err := h.service.FindAllWithCredits()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productsWithCreditsResponse)
}

func (h *productHandler) GetProductWithCredit(c echo.Context) error {
	productWithCreditResponse, err := h.service.FindByIDWithCredit(c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithCreditResponse)
}

func (h *productHandler) CreateProductWithCredit(c echo.Context) error {
	var payload payload.ProductWithCreditPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	payload.Name = strings.TrimSpace(payload.Name)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.Provider = strings.TrimSpace(payload.Provider)

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	productWithCreditResponse, err := h.service.CreateProductWithCredit(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithCreditResponse)
}

func (h *productHandler) UpdateProductWithCredit(c echo.Context) error {
	var payload payload.ProductWithCreditPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	payload.Name = strings.TrimSpace(payload.Name)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.Provider = strings.TrimSpace(payload.Provider)

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	productWithCreditResponse, err := h.service.UpdateProductWithCredit(payload, c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithCreditResponse)
}

func (h *productHandler) DeleteProductWithCredit(c echo.Context) error {
	if err := h.service.DeleteProductWithCredit(c.Param("id")); err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, map[string]any{
		"id":      c.Param("id"),
		"kind":    "product",
		"deleted": true,
	})
}
