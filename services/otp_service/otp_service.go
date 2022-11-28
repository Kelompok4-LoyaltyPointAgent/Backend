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

	_, otpErr := s.otpRepository.FindByPin(pin)
	if otpErr == nil {
		return nil, errors.New("pin already exists")
	}

	otp, err := s.otpRepository.Create(models.OTP{
		UserID:    user.ID,
		Pin:       pin,
		ExpiredAt: time.Now().Add(5 * time.Minute),
		IsUsed:    false,
	})
	if err != nil {
		return nil, err
	}

	// TODO: send pin
	// user, err := s.userRepository.FindByID(userID.String())
	// if err != nil {
	// 	return nil, err
	// }
	// send pin via email to user.Email

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

	if _, err := s.otpRepository.Update(models.OTP{IsUsed: true}, otp.ID); err != nil {
		log.Printf("Error: %s", err)
		return nil, errors.New("internal server error")
	}

	return response.NewVerifyOTPResponse(true), nil
}
