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

	signed, err := token.SignedString([]byte(config.LoadAuthConfig().Secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func GetTokenClaims(c echo.Context) *JWTCustomClaims {
	token := c.Get("token").(*jwt.Token)
	claims := token.Claims.(JWTCustomClaims)

	return &claims
}
