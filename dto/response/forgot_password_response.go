package response

import (
	"time"
)

type RequestForgotPasswordResponse struct {
	ExpiredAt time.Time `json:"expired_at"`
}

type VerifyForgotPasswordResponse struct {
	Verified bool `json:"verified"`
}

func NewRequestForgotPasswordResponse(expiredAt time.Time) *RequestForgotPasswordResponse {
	return &RequestForgotPasswordResponse{
		ExpiredAt: expiredAt,
	}
}

func NewVerifyForgotPasswordResponse(verified bool) *VerifyForgotPasswordResponse {
	return &VerifyForgotPasswordResponse{
		Verified: verified,
	}
}
