package user_handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/services/user_service"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	ChangePassword(c echo.Context) error
	FindUserByID(c echo.Context) error
	Login(c echo.Context) error
	//Admin
	FindAllUser(c echo.Context) error
	FindUserByIDByAdmin(c echo.Context) error
	UpdateUserByAdmin(c echo.Context) error
	DeleteUserByAdmin(c echo.Context) error
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

	user.Name = strings.TrimSpace(user.Name)

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

func (h *userHandler) UpdateUser(c echo.Context) error {
	var userPayload payload.UserPayload

	if err := c.Bind(&userPayload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	userPayload.Name = strings.TrimSpace(userPayload.Name)

	if err := h.validate.Struct(&userPayload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	//Get user id from token
	claims := helper.GetTokenClaims(c)

	user, err := h.service.UpdateProfile(userPayload, claims.ID.String())
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, user)

}

func (h *userHandler) ChangePassword(c echo.Context) error {
	var payload payload.ChangePasswordPayload

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	if err := h.validate.Struct(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	claims := helper.GetTokenClaims(c)

	user, err := h.service.ChangePassword(payload, claims.ID.String())
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	return response.Success(c, "success", http.StatusOK, user)

}

func (h *userHandler) FindUserByID(c echo.Context) error {
	claims := helper.GetTokenClaims(c)

	user, err := h.service.FindByID(claims.ID.String())
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	return response.Success(c, "success", http.StatusOK, user)
}

func (h *userHandler) FindAllUser(c echo.Context) error {
	users, err := h.service.FindAll()
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	return response.Success(c, "success", http.StatusOK, users)
}

func (h *userHandler) FindUserByIDByAdmin(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("id is required"))
	}

	user, err := h.service.FindByID(id)
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	return response.Success(c, "success", http.StatusOK, user)
}

func (h *userHandler) UpdateUserByAdmin(c echo.Context) error {
	var userPayload payload.UserPayload

	if err := c.Bind(&userPayload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	userPayload.Name = strings.TrimSpace(userPayload.Name)

	if err := h.validate.Struct(&userPayload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	id := c.Param("id")

	if id == "" {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("id is required"))
	}

	user, err := h.service.UpdateUserByAdmin(userPayload, id)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, user)

}

func (h *userHandler) DeleteUserByAdmin(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return response.Error(c, "failed", http.StatusBadRequest, errors.New("id is required"))
	}

	_, err := h.service.Delete(id)
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, nil)
}
