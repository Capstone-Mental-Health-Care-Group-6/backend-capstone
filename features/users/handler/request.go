package handler

type RegisterInput struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
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

type UpdateProfile struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RefreshInput struct {
	Token string `json:"access_token" form:"access_token"`
}
