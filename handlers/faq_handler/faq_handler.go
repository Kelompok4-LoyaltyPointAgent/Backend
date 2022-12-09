package faq_handler

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/faq_service"
	"github.com/labstack/echo/v4"
)

type FAQHandler interface {
	GetFAQs(c echo.Context) error
	GetFAQ(c echo.Context) error
	CreateFAQ(c echo.Context) error
	UpdateFAQ(c echo.Context) error
	DeleteFAQ(c echo.Context) error
}

type faqHandler struct {
	validate *validator.Validate
	service  faq_service.FAQService
}

func NewFAQHandler(service faq_service.FAQService) FAQHandler {
	validate := validator.New()
	return &faqHandler{validate, service}
}

func (h *faqHandler) GetFAQs(c echo.Context) error {
	var query string
	var args []any

	category := c.QueryParam("category")
	if category != "" {
		query = "category = ?"
		args = append(args, category)
	}

	faqs, err := h.service.FindAll(query, args)
	if err != nil {
		return response.Error(c, "failed", http.StatusNotFound, err)
	}

	return response.Success(c, "success", http.StatusOK, faqs)
}

func (h *faqHandler) GetFAQ(c echo.Context) error {
	faq, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusNotFound, err)
	}

	return response.Success(c, "success", http.StatusOK, faq)
}

func (h *faqHandler) CreateFAQ(c echo.Context) error {
	var payload payload.FAQPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	payload.Category = strings.TrimSpace(payload.Category)
	payload.Question = strings.TrimSpace(payload.Question)
	payload.Answer = strings.TrimSpace(payload.Answer)

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	faq, err := h.service.Create(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, faq)
}

func (h *faqHandler) UpdateFAQ(c echo.Context) error {
	var payload payload.FAQPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	payload.Category = strings.TrimSpace(payload.Category)
	payload.Question = strings.TrimSpace(payload.Question)
	payload.Answer = strings.TrimSpace(payload.Answer)

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	faq, err := h.service.Update(payload, c.Param("id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusNotFound, err)
	}

	return response.Success(c, "success", http.StatusOK, faq)
}

func (h *faqHandler) DeleteFAQ(c echo.Context) error {
	if err := h.service.Delete(c.Param("id")); err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, nil)
}
