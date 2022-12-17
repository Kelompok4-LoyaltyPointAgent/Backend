package otp_service

import (
	"errors"
	"log"
	"time"

	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/otp_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
)

type OTPService interface {
	CreateOTP(payload payload.RequestOTPPayload) (*response.RequestOTPResponse, error)
	VerifyOTP(payload payload.VerifyOTPPayload) (*response.VerifyOTPResponse, error)
}

type otpService struct {
	otpRepository  otp_repository.OTPRepository
	userRepository user_repository.UserRepository
}

func NewOTPService(otpRepository otp_repository.OTPRepository, userRepository user_repository.UserRepository) OTPService {
	return &otpService{otpRepository, userRepository}
}

func (s *otpService) CreateOTP(payload payload.RequestOTPPayload) (*response.RequestOTPResponse, error) {
	user, err := s.userRepository.FindByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	pin := helper.CreatePin(5)

	if _, err := s.otpRepository.FindByPinAndUserID(pin, user.ID); err == nil {
		return nil, errors.New("failed creating otp")
	}

	otp, err := s.otpRepository.Create(models.OTP{
		UserID:    user.ID,
		Pin:       pin,
		ExpiredAt: time.Now().Add(5 * time.Minute),
	})
	if err != nil {
		return nil, err
	}

	if err := helper.SendOTP(user.Email, helper.OTPEmailData{
		Length: len(pin),
		Pin:    pin,
	}); err != nil {
		return nil, err
	}

	return response.NewRequestOTPResponse(otp.ExpiredAt), nil
}

func (s *otpService) VerifyOTP(payload payload.VerifyOTPPayload) (*response.VerifyOTPResponse, error) {
	user, err := s.userRepository.FindByEmail(payload.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	otp, err := s.otpRepository.FindByPinAndUserID(payload.Pin, user.ID)
	if err != nil {
		return nil, errors.New("invalid pin")
	}

	if time.Now().After(otp.ExpiredAt) {
		return nil, errors.New("otp expired")
	}

	if err := s.otpRepository.Delete(otp.ID); err != nil {
		log.Printf("Error: %s", err)
		return nil, errors.New("internal server error")
	}

	token, err := helper.CreateToken(user.ID, user.Role.String())
	if err != nil {
		return nil, err
	}

	return response.NewVerifyOTPResponse(token), nil
}
