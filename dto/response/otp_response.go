package response

import (
	"time"
)

type RequestOTPResponse struct {
	ExpiredAt time.Time `json:"expired_at"`
}

type VerifyOTPResponse struct {
	Token string `json:"token"`
}

func NewRequestOTPResponse(expiredAt time.Time) *RequestOTPResponse {
	return &RequestOTPResponse{
		ExpiredAt: expiredAt,
	}
}

func NewVerifyOTPResponse(token string) *VerifyOTPResponse {
	return &VerifyOTPResponse{
		Token: token,
	}
}
