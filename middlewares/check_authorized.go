package middlewares

import (
	"errors"
	"net/http"

	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/labstack/echo/v4"
)

func checkRoles(roles []string, userRole string) bool {
	for _, role := range roles {
		if role == userRole {
			return true
		}
	}
	return false
}

func AuthorizedRoles(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := helper.GetTokenClaims(c)

			if !checkRoles(roles, claims.Role) {
				baseResponse := response.Error(c, "Unauthorized", http.StatusUnauthorized, errors.New("unauthorized"))
				return c.JSON(http.StatusUnauthorized, baseResponse)
			}

			return next(c)
		}
	}
}
