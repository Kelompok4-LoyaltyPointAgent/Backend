package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/labstack/echo/v4"
)

func UnauthorizedRole(role []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("token").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			for _, v := range role {
				if claims["role"] == v {
					baseResponse := response.ConvertErrorToBaseResponse("failed", http.StatusUnauthorized, response.EmptyObj{}, "Unauthorized")
					return c.JSON(http.StatusUnauthorized, baseResponse)
				}
			}

			return next(c)
		}
	}
}
