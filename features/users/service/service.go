package service

import (
	"FinalProject/features/users"
	"FinalProject/helper"
	"FinalProject/helper/email"
	"FinalProject/helper/enkrip"
	"errors"
	"strings"
	"time"
)

type UserService struct {
	d     users.UserDataInterface
	j     helper.JWTInterface
	e     enkrip.HashInterface
	email email.EmailInterface
}

func New(data users.UserDataInterface, jwt helper.JWTInterface, email email.EmailInterface, enkrip enkrip.HashInterface) users.UserServiceInterface {
	return &UserService{
		d:     data,
		j:     jwt,
		email: email,
		e:     enkrip,
	}
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	_, err := us.d.GetByEmail(newData.Email)
	if err == nil {
		return nil, errors.New("Email already registered by another user")
	}

	hashPassword, err := us.e.HashPassword(newData.Password)
	if err != nil {
		return nil, errors.New("Hash Password Error")
	}

	newData.Password = hashPassword
	newData.Status = "Active"

	result, err := us.d.Register(newData)

	if err != nil {
		return nil, errors.New("Failed to Register")
	}
	return result, nil
}

func (us *UserService) Login(email, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "Incorrect Password") {
			return nil, errors.New("Incorrect Password")
		}
		if strings.Contains(err.Error(), "Not Found") {
			return nil, errors.New("User Not Found / User Inactive")
		}
		return nil, errors.New("Process Failed")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("Token Process Failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func (us *UserService) GenerateJwt(email string) (*users.UserCredential, error) {
	result, err := us.d.GetByEmail(email)
	if err != nil {
		return nil, errors.New("Process Failed")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Role, result.Status)

	if tokenData == nil {
		return nil, errors.New("Token Process Failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func (us *UserService) TokenResetVerify(code string) (*users.UserResetPass, error) {
	result, err := us.d.GetByCode(code)
	if err != nil {
		return nil, errors.New("Failed to verify token")
	}

	if result.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return result, nil
}

func (us *UserService) ForgetPasswordWeb(email string) error {

	user, err := us.d.GetByEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errors.New("data not found")
		}
		return errors.New("process failed")
	}

	email = user.Email

	header, htmlBody, code := us.email.HTMLBody(user.Role, user.Name)

	if err := us.d.InsertCode(email, code); err != nil {
		return errors.New("Insert Code Failed")
	}

	err = us.email.SendEmail(email, header, htmlBody)

	if err != nil {
		return errors.New("Send Email Error")
	}

	return nil
}

func (us *UserService) ResetPassword(code, email, password string) error {
	hashPassword, err := us.e.HashPassword(password)
	if err != nil {
		return errors.New("Hash Password Error")
	}
	password = hashPassword

	if err := us.d.ResetPassword(code, email, password); err != nil {
		return errors.New("Reset Password Process Failed")
	}

	return nil
}

func (us *UserService) UpdateProfile(id int, newData users.UpdateProfile) (bool, error) {
	hashPassword, err := us.e.HashPassword(newData.Password)
	if err != nil {
		return false, errors.New("Hash Password Error")
	}
	newData.Password = hashPassword
	result, err := us.d.UpdateProfile(id, newData)
	if err != nil {
		return false, errors.New("Update Process Failed")
	}
	return result, nil
}
