package forgot_password_service

import (
	"errors"
	"log"
	"time"

	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/forgot_password_repository"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
)

type ForgotPasswordService interface {
	RequestForgotPassword(payload payload.RequestForgotPasswordPayload) error
	SubmitForgotPassword(payload payload.SubmitForgotPasswordPayload) error
}

type forgotPasswordService struct {
	forgotPasswordRepository forgot_password_repository.ForgotPasswordRepository
	userRepository           user_repository.UserRepository
}

func NewForgotPasswordService(
	forgotPasswordRepository forgot_password_repository.ForgotPasswordRepository,
	userRepository user_repository.UserRepository,
) ForgotPasswordService {
	return &forgotPasswordService{forgotPasswordRepository, userRepository}
}

func (s *forgotPasswordService) RequestForgotPassword(payload payload.RequestForgotPasswordPayload) error {
	user, err := s.userRepository.FindByEmail(payload.Email)
	if err != nil {
		return err
	}

	accessKey := helper.CreateAccessKey(50)

	if _, err := s.forgotPasswordRepository.Create(models.ForgotPassword{
		UserID:    user.ID,
		AccessKey: accessKey,
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}); err != nil {
		return err
	}

	if err := helper.SendAccessKey(user.Email, helper.AccessKeyEmailData{
		AccessKey: accessKey,
		FormURL:   "http://localhost/reset-password/" + accessKey,
	}); err != nil {
		return err
	}

	return nil
}

func (s *forgotPasswordService) SubmitForgotPassword(payload payload.SubmitForgotPasswordPayload) error {
	forgotPassword, err := s.forgotPasswordRepository.FindByAccessKey(payload.AccessKey)
	if err != nil {
		return errors.New("invalid access key")
	}

	if time.Now().After(forgotPassword.ExpiredAt) {
		return errors.New("access key expired")
	}

	user, err := s.userRepository.FindByID(forgotPassword.UserID.String())
	if err != nil {
		return err
	}

	updates := models.User{}
	if err := updates.HashPassword(payload.NewPassword); err != nil {
		return err
	}
	if _, err := s.userRepository.Update(updates, user.ID.String()); err != nil {
		return err
	}

	if err := s.forgotPasswordRepository.Delete(forgotPassword.ID); err != nil {
		log.Printf("Error: %s", err)
		return errors.New("internal server error")
	}

	return nil
}
