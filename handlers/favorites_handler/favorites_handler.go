package favorites_handler

import (
	"net/http"

	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/services/favorites_service"
	"github.com/labstack/echo/v4"
)

type FavoritesHandler interface {
	FindAll(c echo.Context) error
	Create(c echo.Context) error
	Delete(c echo.Context) error
}

type favoritesHandler struct {
	service favorites_service.FavoritesService
}

func NewFavoritesHandler(service favorites_service.FavoritesService) FavoritesHandler {
	return &favoritesHandler{service}
}

func (h *favoritesHandler) FindAll(c echo.Context) error {
	claims := helper.GetTokenClaims(c)
	favorites, err := h.service.FindAll(claims)
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}
	return response.Success(c, "success", http.StatusOK, favorites)
}

func (h *favoritesHandler) Create(c echo.Context) error {
	var payload payload.FavoritesPayload
	if err := c.Bind(&payload); err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}

	claims := helper.GetTokenClaims(c)
	favorites, err := h.service.Create(payload, claims.ID.String())
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}
	return response.Success(c, "success", http.StatusOK, favorites)
}

func (h *favoritesHandler) Delete(c echo.Context) error {
	claims := helper.GetTokenClaims(c)
	err := h.service.Delete(claims.ID.String(), c.Param("product_id"))
	if err != nil {
		return response.Error(c, "failed", http.StatusBadRequest, err)
	}
	return response.Success(c, "success", http.StatusOK, nil)
}
