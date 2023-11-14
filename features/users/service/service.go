package service

import (
	"FinalProject/features/users"
	"FinalProject/helper"
	"errors"
	"strings"
)

type UserService struct {
	d users.UserDataInterface
	j helper.JWTInterface
}

func New(data users.UserDataInterface, jwt helper.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		j: jwt,
	}
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	result, err := us.d.Register(newData)
	if err != nil {
		return nil, errors.New("Register Process Failed")
	}
	return result, nil
}

func (us *UserService) Login(email, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("process failed")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}
