package payload

type UserPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Points   int    `json:"points" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	OldPassword     string `json:"old_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}