package handlerimport

import (
	"mime/multipart"
)

type PatientRequest struct {
	UserID         uint           `json:"user_id" form:"user_id"`
	Name           string         `json:"name" form:"name"`
	DateOfBirth    string         `json:"date_of_birth" form:"date_of_birth"`
	PlaceOfBirth   string         `json:"place_of_birth" form:"place_of_birth"`
	Gender         string         `json:"gender" form:"gender"`
	MarriageStatus string         `json:"marriage_status" form:"marriage_status"`
	Avatar         multipart.File `json:"avatar" form:"avatar"`
	Address        string         `json:"address" form:"address"`
}
