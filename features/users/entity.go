package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

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

type UserResetPass struct {
	Email     string
	Code      string
	ExpiresAt time.Time
}

type UpdateProfile struct {
	Name     string
	Email    string
	Password string
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	LoginGoogle() echo.HandlerFunc
	CallbackGoogle() echo.HandlerFunc
	ForgetPasswordWeb() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
	ForgetPasswordVerify() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	RefreshToken() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*UserCredential, error)
	GenerateJwt(email string) (*UserCredential, error)
	ForgetPasswordWeb(email string) error
	TokenResetVerify(code string) (*UserResetPass, error)
	ResetPassword(code, email string, password string) error
	UpdateProfile(id int, newData UpdateProfile) (bool, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*User, error)
	GetByID(id int) (User, error)
	GetByEmail(email string) (*User, error)
	InsertCode(email string, code string) error
	DeleteCode(email string) error
	GetByCode(code string) (*UserResetPass, error)
	ResetPassword(code, email string, password string) error
	UpdateProfile(id int, newData UpdateProfile) (bool, error)
}
