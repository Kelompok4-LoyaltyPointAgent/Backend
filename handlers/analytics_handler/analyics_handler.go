package analytics_handler

import (
	"net/http"

	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/services/analytics_service"
	"github.com/labstack/echo/v4"
)

type AnalyticsHandler interface {
	Analytics(c echo.Context) error
	DataForManageStockAdmin(c echo.Context) error
}

type analyticsHandler struct {
	service analytics_service.AnalyticsService
}

func NewAnalyticsHandler(service analytics_service.AnalyticsService) AnalyticsHandler {
	return &analyticsHandler{service}
}

func (h *analyticsHandler) Analytics(c echo.Context) error {
	analytics, err := h.service.Analytics()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, analytics)
}

func (h *analyticsHandler) DataForManageStockAdmin(c echo.Context) error {
	data, err := h.service.DataForManageStockAdmin()
	if err != nil {
		return response.Error(c, "failed", http.StatusInternalServerError, err)
	}

	return response.Success(c, "success", http.StatusOK, data)
}
