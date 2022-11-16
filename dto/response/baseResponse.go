package dto

import "strings"

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
