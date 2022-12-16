package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/kelompok4-loyaltypointagent/backend/config"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
	jwt.StandardClaims
}

func CreateToken(id uuid.UUID, role string) (string, error) {
	exp := time.Duration(config.LoadAuthConfig().ExpHours) * time.Hour
	claims := JWTCustomClaims{
		id,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.LoadAuthConfig().Secret))
}

func GetTokenClaims(c echo.Context) *JWTCustomClaims {
	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(*JWTCustomClaims)

	return claims
}
