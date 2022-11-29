package user_service

import (
	"errors"

	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
)

type UserService interface {
	FindByID(id string) (response.UserResponse, error)
	FindByIDByAdmin(id string) (models.User, error)
	FindByEmail(email string) (response.UserResponse, error)
	FindAll() ([]response.UserResponse, error)
	Create(payload payload.UserPayload) (response.UserResponse, error)
	Login(payload payload.LoginPayload) (response.LoginResponse, error)
	UpdateProfile(payload payload.UserPayload, id string) (response.UserResponse, error)
	UpdateUserByAdmin(payload payload.UserPayload, id string) (response.UserResponse, error)
	ChangePassword(payload payload.ChangePasswordPayload, id string) (response.UserResponse, error)
	Delete(id string) (response.UserResponse, error)
}

type userService struct {
	repository user_repository.UserRepository
}

func NewUserService(repository user_repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) FindByID(id string) (response.UserResponse, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}

func (s *userService) FindByIDByAdmin(id string) (models.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *userService) Create(payload payload.UserPayload) (response.UserResponse, error) {
	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     "User",
		Points:   payload.Points,
	}

	if err := user.HashPassword(user.Password); err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.Create(user)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}

func (s *userService) FindAll() ([]response.UserResponse, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return []response.UserResponse{}, err
	}

	var usersResponses []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			ID:     user.ID.String(),
			Name:   user.Name,
			Email:  user.Email,
			Points: user.Points,
		}
		usersResponses = append(usersResponses, userResponse)
	}
	return usersResponses, nil
}

func (s *userService) UpdateProfile(payload payload.UserPayload, id string) (response.UserResponse, error) {

	modelUser := models.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	_, err := s.repository.Update(modelUser, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}

func (s *userService) UpdateUserByAdmin(payload payload.UserPayload, id string) (response.UserResponse, error) {

	userModel := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Points:   payload.Points,
	}

	if payload.Password != "" {
		if err := userModel.HashPassword(userModel.Password); err != nil {
			return response.UserResponse{}, err
		}
	}

	_, err := s.repository.Update(userModel, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}

func (s *userService) Delete(id string) (response.UserResponse, error) {
	user, err := s.repository.Delete(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}

func (s *userService) FindByEmail(email string) (response.UserResponse, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}

func (s *userService) Login(payload payload.LoginPayload) (response.LoginResponse, error) {
	user, err := s.repository.FindByEmail(payload.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		return response.LoginResponse{}, err
	}

	token, err := helper.CreateToken(user.ID, user.Role)
	if err != nil {
		return response.LoginResponse{}, err
	}

	loginResponse := response.LoginResponse{
		Token: token,
	}
	return loginResponse, nil
}

func (s *userService) ChangePassword(payload payload.ChangePasswordPayload, id string) (response.UserResponse, error) {

	//Get User by ID
	getUser, err := s.repository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	if err := getUser.CheckPassword(payload.OldPassword); err != nil {
		return response.UserResponse{}, err
	}

	if payload.NewPassword != payload.ConfirmPassword {
		return response.UserResponse{}, errors.New("New Password and Confirm Password not match")
	}

	getUser.Password = payload.NewPassword

	if err := getUser.HashPassword(getUser.Password); err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.Update(getUser, id)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
	return userResponse, nil
}
