package payload

type UserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
