package handler

type RegisterInput struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ForgetPasswordInput struct {
	Email string `json:"email" form:"email"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" form:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required"`
}
