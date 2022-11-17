package hello_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloHandler interface {
	Greeting(c echo.Context) error
}

type helloHandler struct{}

func NewHelloHandler() HelloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Greeting(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})
}
