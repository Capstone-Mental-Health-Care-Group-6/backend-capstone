package users

import "github.com/labstack/echo/v4"

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

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*UserCredential, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email string, password string) (*User, error)
}
