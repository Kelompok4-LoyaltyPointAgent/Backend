package payload

type UserPayload struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Points   uint   `json:"points"`
}

type UserPayloadByAdmin struct {
	Name   string `json:"name" validate:"required,min=1,max=100"`
	Email  string `json:"email" validate:"required,email"`
	Points uint   `json:"points"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	OldPassword     string `json:"old_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

type ChangePasswordFromResetPasswordPayload struct {
	NewPassword     string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

type CheckPasswordPayload struct {
	CheckPassword string `json:"check_password" validate:"requred,min=8"`
}
