package user_handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/user_service"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	validate *validator.Validate
	service  user_service.UserService
}

func NewUserHandler(service user_service.UserService) UserHandler {
	validate := validator.New()
	return &userHandler{validate, service}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	var user payload.UserPayload

	if err := c.Bind(&user); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&user); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	userResponse, err := h.service.Create(user)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, userResponse)
}

func (h *userHandler) Login(c echo.Context) error {
	var payload payload.LoginPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	loginResponse, err := h.service.Login(payload)
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("invalid email or password"))
	}

	return response.Success(c, "success", http.StatusOK, loginResponse)
}
