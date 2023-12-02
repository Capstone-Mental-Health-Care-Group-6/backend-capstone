package users

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type UserCredential struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Access map[string]any `json:"token"`
}

type UserInfo struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type UserDashboard struct {
	TotalUser         int `json:"total_user"`
	TotalUserBaru     int `json:"total_user_baru"`
	TotalUserActive   int `json:"total_user_active"`
	TotalUserInactive int `json:"total_user_inactive"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	LoginGoogle() echo.HandlerFunc
	CallbackGoogle() echo.HandlerFunc
	GetUsers() echo.HandlerFunc
	UserDashboard() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*UserCredential, error)
	GenerateJwt(email string) (*UserCredential, error)
	GetUsers(status, name string) ([]User, error)
	UserDashboard() (UserDashboard, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetUsers(status, name string) ([]User, error)
	UserDashboard() (UserDashboard, error)
}
