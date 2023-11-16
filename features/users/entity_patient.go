package users

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"
)

type Patiententity struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	UserID         uint   `json:"user_id"`
	DateOfBirth    string `json:"date_of_birth"`
	PlaceOfBirth   string `json:"place_of_birth"`
	Gender         string `json:"gender"`
	MarriageStatus string `json:"marriage_status"`
	Avatar         string `json:"avatar"`
	Address        string `json:"address"`
}

type AvatarPhoto struct {
	Avatar multipart.File `json:"avatar"`
}

type PatientHandlerInterface interface {
	GetPatients() echo.HandlerFunc
	GetPatient() echo.HandlerFunc
	CreatePatient() echo.HandlerFunc
}

type PatientServiceInterface interface {
	GetPatients() ([]Patiententity, error)
	GetPatient(id int) ([]Patiententity, error)
	CreatePatient(newData Patiententity) (*Patiententity, error)
	PhotoUpload(newData AvatarPhoto) (string, error)
}

type PatientDataInterface interface {
	GetAll() ([]Patiententity, error)
	GetByID(id int) ([]Patiententity, error)
	Insert(newData Patiententity) (*Patiententity, error)
}
