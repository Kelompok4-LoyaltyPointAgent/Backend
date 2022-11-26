package response

import (
	"time"
)

type RequestOTPResponse struct {
	ExpiredAt time.Time `json:"expired_at"`
}

type VerifyOTPResponse struct {
	Verified bool `json:"verified"`
}

func NewRequestOTPResponse(expiredAt time.Time) *RequestOTPResponse {
	return &RequestOTPResponse{
		ExpiredAt: expiredAt,
	}
}

func NewVerifyOTPResponse(verified bool) *VerifyOTPResponse {
	return &VerifyOTPResponse{
		Verified: verified,
	}
}
