package handler

import (
	"mime/multipart"
)

type PatientRequest struct {
	Name        string         `json:"name" form:"name" validate:"required"`
	Email       string         `json:"email" form:"email" validate:"required"`
	Password    string         `json:"password" form:"password" validate:"required"`
	DateOfBirth string         `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Gender      string         `json:"gender" form:"gender" validate:"required"`
	Avatar      multipart.File `json:"avatar" form:"avatar" validate:"required"`
	Phone       string         `json:"phone_number" form:"phone_number" validate:"required"`
	Role        string         `json:"role" form:"role" validate:"required"`
	Status      string         `json:"status" form:"status" validate:"required"`
}

type UpdateProfile struct {
	Name        string         `json:"name" form:"name"`
	Email       string         `json:"email" form:"email"`
	DateOfBirth string         `json:"date_of_birth" form:"date_of_birth"`
	Gender      string         `json:"gender" form:"gender"`
	Phone       string         `json:"phone" form:"phone"`
	Avatar      multipart.File `json:"avatar" form:"avatar"`
}

type UpdatePassword struct {
	Password string `json:"password" form:"password"`
}

type LoginPatient struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
