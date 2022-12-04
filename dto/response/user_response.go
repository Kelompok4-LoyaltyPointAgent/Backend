package response

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Points uint   `json:"points"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
