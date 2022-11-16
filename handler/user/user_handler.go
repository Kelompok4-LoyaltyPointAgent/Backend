package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	userService "github.com/kelompok4-loyaltypointagent/backend/services/user"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

type UserHandler interface {
	CreateUser(c echo.Context) error
}

type userHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *userHandler {
	validate = validator.New()
	return &userHandler{service}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	var user payload.UserPayload

	if err := c.Bind(&user); err != nil {
		baseResponse := response.ConvertErrorToBaseResponse("failed", http.StatusBadRequest, response.EmptyObj{}, err.Error())
		return c.JSON(http.StatusBadRequest, baseResponse)
	}

	if err := validate.Struct(&user); err != nil {
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
