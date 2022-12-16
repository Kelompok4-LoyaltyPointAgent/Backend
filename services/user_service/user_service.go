package user_service

import (
	"errors"

	"github.com/kelompok4-loyaltypointagent/backend/constant"
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
	FindAll(filter string) ([]response.UserResponse, error)
	Create(payload payload.UserPayload) (response.UserResponse, error)
	Login(payload payload.LoginPayload) (response.LoginResponse, error)
	UpdateProfile(payload payload.UserPayload, id string) (response.UserResponse, error)
	UpdateUserByAdmin(payload payload.UserPayloadByAdmin, id string) (response.UserResponse, error)
	ChangePassword(payload payload.ChangePasswordPayload, id string) (response.UserResponse, error)
	ChangePasswordFromResetPassword(payload payload.ChangePasswordFromResetPasswordPayload, id string) (response.UserResponse, error)
	CheckPassword(payload payload.CheckPasswordPayload, id string) (bool, error)
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
	return *response.NewUserResponse(user), nil
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
		Role:     constant.UserRoleUser,
		Points:   payload.Points,
	}

	if err := user.HashPassword(user.Password); err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.Create(user)
	if err != nil {
		return response.UserResponse{}, err
	}

	return *response.NewUserResponse(user), nil
}

func (s *userService) FindAll(filter string) ([]response.UserResponse, error) {

	var query string
	var args string

	if filter == "" {
		query = ""
	} else if filter == constant.UserRoleAdmin.String() {
		query = "role = ?"
		args = constant.UserRoleAdmin.String()
	} else if filter == constant.UserRoleUser.String() {
		query = "role = ?"
		args = constant.UserRoleUser.String()
	} else {
		return []response.UserResponse{}, errors.New("role not found")
	}

	users, err := s.repository.FindAll(query, args)
	if err != nil {
		return []response.UserResponse{}, err
	}

	return *response.NewUsersResponse(users), nil
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

	return *response.NewUserResponse(user), nil
}

func (s *userService) UpdateUserByAdmin(payload payload.UserPayloadByAdmin, id string) (response.UserResponse, error) {

	userModel := models.User{
		Name:   payload.Name,
		Email:  payload.Email,
		Points: payload.Points,
	}

	_, err := s.repository.Update(userModel, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	return *response.NewUserResponse(user), nil
}

func (s *userService) Delete(id string) (response.UserResponse, error) {
	user, err := s.repository.Delete(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	return *response.NewUserResponse(user), nil
}

func (s *userService) FindByEmail(email string) (response.UserResponse, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return response.UserResponse{}, err
	}
	return *response.NewUserResponse(user), nil
}

func (s *userService) Login(payload payload.LoginPayload) (response.LoginResponse, error) {
	user, err := s.repository.FindByEmail(payload.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	if err := user.CheckPassword(payload.Password); err != nil {
		return response.LoginResponse{}, err
	}

	token, err := helper.CreateToken(user.ID, user.Role.String())
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
		return response.UserResponse{}, errors.New("new Password and confirm Password not match")
	}

	getUser.Password = payload.NewPassword

	if err := getUser.HashPassword(getUser.Password); err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.Update(getUser, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	return *response.NewUserResponse(user), nil
}

func (s *userService) ChangePasswordFromResetPassword(payload payload.ChangePasswordFromResetPasswordPayload, id string) (response.UserResponse, error) {

	//Get User by ID
	getUser, err := s.repository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	if payload.NewPassword != payload.ConfirmPassword {
		return response.UserResponse{}, errors.New("new Password and confirm Password not match")
	}

	getUser.Password = payload.NewPassword

	if err := getUser.HashPassword(getUser.Password); err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.Update(getUser, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	return *response.NewUserResponse(user), nil

}

func (s *userService) CheckPassword(payload payload.CheckPasswordPayload, id string) (bool, error) {

	getUser, err := s.repository.FindByID(id)
	if err != nil {
		return false, err
	}

	if err := getUser.CheckPassword(payload.CheckPassword); err != nil {
		return false, err
	}

	return true, nil
}
