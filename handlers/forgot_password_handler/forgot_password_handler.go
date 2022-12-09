package forgot_password_handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/forgot_password_service"
	"github.com/labstack/echo/v4"
)

type ForgotPasswordHandler interface {
	RequestForgotPassword(c echo.Context) error
	SubmitForgotPassword(c echo.Context) error
}

type forgot_passwordHandler struct {
	validate *validator.Validate
	service  forgot_password_service.ForgotPasswordService
}

func NewForgotPasswordHandler(service forgot_password_service.ForgotPasswordService) ForgotPasswordHandler {
	validate := validator.New()
	return &forgot_passwordHandler{validate, service}
}

func (h *forgot_passwordHandler) RequestForgotPassword(c echo.Context) error {
	var payload payload.RequestForgotPasswordPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.service.RequestForgotPassword(payload); err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, nil)
}

func (h *forgot_passwordHandler) SubmitForgotPassword(c echo.Context) error {
	var payload payload.SubmitForgotPasswordPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if payload.NewPassword != payload.ConfirmPassword {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("invalid password"))
	}

	if err := h.service.SubmitForgotPassword(payload); err != nil {
		if err.Error() == "internal server error" {
			return response.Error(c, "failed", http.StatusInternalServerError, err)
		}
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	return response.Success(c, "success", http.StatusOK, nil)
}
