package product_handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/product_service"
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	GetProducts(c echo.Context) error
	CreateProduct(c echo.Context) error
}

type productHandler struct {
	validate *validator.Validate
	service  product_service.ProductService
}

func NewProductHandler(service product_service.ProductService) ProductHandler {
	validate := validator.New()
	return &productHandler{validate, service}
}

func (h *productHandler) GetProducts(c echo.Context) error {
	productsResponse, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productsResponse)
}

func (h *productHandler) CreateProduct(c echo.Context) error {
	var payload payload.ProductPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	productResponse, err := h.service.CreateProduct(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productResponse)
}
