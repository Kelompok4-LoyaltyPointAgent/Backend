package response

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors"`
	Status  int         `json:"status"`
}

type EmptyObj struct{}

func ConvertToBaseResponse(message string, status int, data interface{}) BaseResponse {
	return BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  nil,
	}
}

func ConvertErrorToBaseResponse(message string, status int, data interface{}, err string) BaseResponse {
	splittedError := strings.Split(err, "\n")
	return BaseResponse{
		Message: message,
		Data:    data,
		Status:  status,
		Errors:  splittedError,
	}
}

func Success(ctx echo.Context, message string, status int, data interface{}) error {
	response := ConvertToBaseResponse(message, status, data)
	return ctx.JSON(status, response)
}

func Error(ctx echo.Context, message string, status int, err error) error {
	response := ConvertErrorToBaseResponse(message, status, EmptyObj{}, err.Error())
	return ctx.JSON(status, response)
}
