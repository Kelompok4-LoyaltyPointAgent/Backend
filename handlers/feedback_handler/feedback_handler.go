package feedback_handler

import (
	"net/http"

	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/services/feedback_service"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler interface {
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Create(c echo.Context) error
}

type feedbackHandler struct {
	service feedback_service.FeedbackService
}

func NewFeedbackHandler(service feedback_service.FeedbackService) FeedbackHandler {
	return &feedbackHandler{service}
}

func (h *feedbackHandler) Create(c echo.Context) error {
	var payload payload.FeedbackPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	//Claims ID
	claims := helper.GetTokenClaims(c)

	feedback, err := h.service.Create(payload, claims.ID.String())
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, feedback)

}

func (h *feedbackHandler) FindAll(c echo.Context) error {
	feedback, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, feedback)
}

func (h *feedbackHandler) FindByID(c echo.Context) error {
	id := c.Param("id")

	feedback, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, feedback)
}
