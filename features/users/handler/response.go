package handler

type RegisterResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
}

type LoginResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token any    `json:"token"`
}

type UserInfo struct {
	Name   string `json:"name" form:"name"`
	Email  string `json:"email" form:"email"`
	Status string `json:"status" form:"status"`
}

type DashboardResponse struct {
	TotalUser         int `json:"total_user"`
	TotalUserBaru     int `json:"total_user_baru"`
	TotalUserActive   int `json:"total_user_active"`
	TotalUserInactive int `json:"total_user_inactive"`
}
