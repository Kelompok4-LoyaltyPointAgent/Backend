package user

import (
	"github.com/kelompok4-loyaltypointagent/backend/dto/payload"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	userModel "github.com/kelompok4-loyaltypointagent/backend/models/user"
	userRepo "github.com/kelompok4-loyaltypointagent/backend/repositories/user"
)

type UserService interface {
	FindByID(id string) (response.UserResponse, error)
	Create(payload payload.UserPayload) (response.UserResponse, error)
	FindAll() ([]response.UserResponse, error)
	UpdateProfile(payload payload.UserPayload, id string) (response.UserResponse, error)
	Delete(id string) (response.UserResponse, error)
	FindByEmail(email string) (response.UserResponse, error)
}

type userService struct {
	repository userRepo.UserRepository
}

func NewUserService(repository userRepo.UserRepository) *userService {
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
	user := userModel.User{
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
	userUpdate := userModel.User{
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
