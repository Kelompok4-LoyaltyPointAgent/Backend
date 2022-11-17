package user_handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/user_service"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
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
		baseResponse := response.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, response.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if err := h.validate.Struct(&user); err != nil {
		baseResponse := response.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, response.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	userResponse, err := h.service.Create(user)
	if err != nil {
		baseResponse := response.ConvertErrorToBaseResponse("failed", http.StatusInternalServerError, response.EmptyObj{}, err.Error())
		return c.JSON(http.StatusInternalServerError, baseResponse)
	}

	baseResponse := response.ConvertToBaseResponse("success", http.StatusOK, userResponse)
	return c.JSON(http.StatusOK, baseResponse)
}
