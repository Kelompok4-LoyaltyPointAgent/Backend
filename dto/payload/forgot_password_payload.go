package payload

type RequestForgotPasswordPayload struct {
	Email string `json:"email" validate:"required,email"`
}

type SubmitForgotPasswordPayload struct {
	AccessKey       string `json:"access_key" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
