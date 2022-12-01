package product_handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/product_service"
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	//Product With Credits
	GetProductsWithCredits(c echo.Context) error
	GetProductWithCredit(c echo.Context) error
	GetProductByProviderWithCredits(provider string, c echo.Context) error
	GetProductByRecommendedWithCredits(c echo.Context) error
	CreateProductWithCredit(c echo.Context) error
	UpdateProductWithCredit(c echo.Context) error
	DeleteProductWithCredit(c echo.Context) error

	//Product With Packages
	GetProductsWithPackages(c echo.Context) error
	GetProductWithPackage(c echo.Context) error
	GetProductByProviderWithPackages(provider string, c echo.Context) error
	GetProductByRecommendedWithPackages(c echo.Context) error
	CreateProductWithPackage(c echo.Context) error
	UpdateProductWithPackage(c echo.Context) error
	DeleteProductWithPackage(c echo.Context) error
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
	provider := c.QueryParam("provider")
	recommended := c.QueryParam("recommended")

	if provider != "" && recommended != "" {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("provider and recommended cannot be used at the same time"))
	} else if recommended == "true" {
		return h.GetProductByRecommendedWithCredits(c)
	} else if provider != "" {
		return h.GetProductByProviderWithCredits(provider, c)
	}

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

func (h *productHandler) GetProductByProviderWithCredits(provider string, c echo.Context) error {
	productsWithPackagesResponse, err := h.service.FindByProviderWithCredit(provider)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productsWithPackagesResponse)
}

func (h *productHandler) GetProductByRecommendedWithCredits(c echo.Context) error {
	productsWithPackagesResponse, err := h.service.FindByRecommendedWithCredit()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productsWithPackagesResponse)
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

	file, err := c.FormFile("product_picture")
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	payload.ProductPicture = file

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

	file, err := c.FormFile("product_picture")
	if err == nil {
		payload.ProductPicture = file
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
		"kind":    "Credit",
		"deleted": true,
	})
}

func (h *productHandler) GetProductsWithPackages(c echo.Context) error {
	provider := c.QueryParam("provider")
	recommended := c.QueryParam("recommended")

	if provider != "" && recommended != "" {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("provider and recommended cannot be used at the same time"))
	} else if recommended == "true" {
		return h.GetProductByRecommendedWithPackages(c)
	} else if provider != "" {
		return h.GetProductByProviderWithPackages(provider, c)
	}

	productsWithPackagesResponse, err := h.service.FindAllWithPackages()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productsWithPackagesResponse)
}

func (h *productHandler) GetProductWithPackage(c echo.Context) error {
	productWithPackageResponse, err := h.service.FindByIDWithPackages(c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithPackageResponse)
}

func (h *productHandler) GetProductByProviderWithPackages(provider string, c echo.Context) error {
	productWithPackageResponse, err := h.service.FindByProviderWithPackages(provider)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithPackageResponse)
}

func (h *productHandler) GetProductByRecommendedWithPackages(c echo.Context) error {
	productWithPackageResponse, err := h.service.FindByRecommendedWithPackages()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithPackageResponse)
}

func (h *productHandler) CreateProductWithPackage(c echo.Context) error {
	var payload payload.ProductWithPackagesPayload

	if err := c.Bind(&payload); err != nil {
		log.Println(err)
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	file, err := c.FormFile("product_picture")
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	payload.ProductPicture = file

	productWithPackageResponse, err := h.service.CreateProductWithPackages(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithPackageResponse)
}

func (h *productHandler) UpdateProductWithPackage(c echo.Context) error {
	var payload payload.ProductWithPackagesPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	file, err := c.FormFile("product_picture")
	if err == nil {
		payload.ProductPicture = file
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	productWithPackageResponse, err := h.service.UpdateProductWithPackages(payload, c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, productWithPackageResponse)
}

func (h *productHandler) DeleteProductWithPackage(c echo.Context) error {
	if err := h.service.DeleteProductWithPackages(c.Param("id")); err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, map[string]any{
		"id":      c.Param("id"),
		"kind":    "Package",
		"deleted": true,
	})
}
