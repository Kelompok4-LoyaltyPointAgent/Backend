package response

import "github.com/kelompok4-loyaltypointagent/backend/models"

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Points int    `json:"points"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewUserResponse(user models.User) *UserResponse {
	return &UserResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Points: user.Points,
	}
}
