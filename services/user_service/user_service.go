package user_service

import (
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/user_repository"
)

type UserService interface {
	FindByID(id string) (response.UserResponse, error)
	Create(payload payload.UserPayload) (response.UserResponse, error)
	FindAll() ([]response.UserResponse, error)
	UpdateProfile(payload payload.UserPayload, id string) (response.UserResponse, error)
	Delete(id string) (response.UserResponse, error)
	FindByEmail(email string) (response.UserResponse, error)
	Login(email, password string) (response.LoginResponse, error)
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
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return userResponse, nil
}

func (s *userService) Create(payload payload.UserPayload) (response.UserResponse, error) {
	user := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     "User",
	}

	if err := user.HashPassword(user.Password); err != nil {
		return response.UserResponse{}, err
	}

	user, err := s.repository.Create(user)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return userResponse, nil
}

func (s *userService) FindAll() ([]response.UserResponse, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return []response.UserResponse{}, err
	}
	var usersResponse []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		}
		usersResponse = append(usersResponse, userResponse)
	}
	return usersResponse, nil
}

func (s *userService) UpdateProfile(payload payload.UserPayload, id string) (response.UserResponse, error) {
	userUpdate := models.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	user, err := s.repository.Update(userUpdate, id)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return userResponse, nil
}

func (s *userService) Delete(id string) (response.UserResponse, error) {
	user, err := s.repository.Delete(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return userResponse, nil
}

func (s *userService) FindByEmail(email string) (response.UserResponse, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
	return userResponse, nil
}

func (s *userService) Login(email, password string) (response.LoginResponse, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	if err := user.CheckPassword(password); err != nil {
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
