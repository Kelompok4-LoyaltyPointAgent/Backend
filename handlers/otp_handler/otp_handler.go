package otp_handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/otp_service"
	"github.com/labstack/echo/v4"
)

type OTPHandler interface {
	RequestOTP(c echo.Context) error
	VerifyOTP(c echo.Context) error
}

type otpHandler struct {
	validate *validator.Validate
	service  otp_service.OTPService
}

func NewOTPHandler(service otp_service.OTPService) OTPHandler {
	validate := validator.New()
	return &otpHandler{validate, service}
}

func (h *otpHandler) RequestOTP(c echo.Context) error {
	var payload payload.RequestOTPPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	otp, err := h.service.CreateOTP(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, otp)
}

func (h *otpHandler) VerifyOTP(c echo.Context) error {
	var payload payload.VerifyOTPPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	otp, err := h.service.VerifyOTP(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	return response.Success(c, "success", http.StatusOK, otp)
}
