package handler

type RegisterResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type LoginResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token any    `json:"token"`
}
