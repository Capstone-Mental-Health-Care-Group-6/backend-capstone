package handler

import "mime/multipart"

type InputRequest struct {
	Name         string `json:"name" form:"name" validate:"required"`
	Sessions     uint   `json:"sessions" form:"sessions" validate:"required"`
	Type         string `json:"type" form:"type" validate:"required"`
	Price        uint   `json:"price" form:"price" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	ActivePriode uint   `json:"active_priode" form:"active_priode" validate:"required"`
}

type InputFileRequest struct {
	Avatar multipart.File `json:"avatar" form:"avatar" validate:"required"`
}
