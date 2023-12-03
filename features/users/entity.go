package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role     string
	Status   string
}

type UserCredential struct {
	Name   string
	Email  string
	Access map[string]any
}

type UserResetPass struct {
	Email     string
	Code      string
	ExpiresAt time.Time
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	LoginGoogle() echo.HandlerFunc
	CallbackGoogle() echo.HandlerFunc
	ForgetPasswordWeb() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
	ForgetPasswordVerify() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*UserCredential, error)
	GenerateJwt(email string) (*UserCredential, error)
	ForgetPasswordWeb(email string) error
	TokenResetVerify(code string) (*UserResetPass, error)
	ResetPassword(code, email string, password string) error
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*User, error)
	GetByEmail(email string) (*User, error)
	InsertCode(email string, code string) error
	DeleteCode(email string) error
	GetByCode(code string) (*UserResetPass, error)
	ResetPassword(code, email string, password string) error
}
