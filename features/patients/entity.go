package patients

import (
	"github.com/labstack/echo/v4"
	"mime/multipart"
)

type Patiententity struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	DateOfBirth string
	Gender      string
	Avatar      string
	Phone       string
	Role        string
	Status      string
}

type Patientdetail struct {
	ID          uint
	Name        string
	Email       string
	DateOfBirth string
	Gender      string
	Avatar      string
	Phone       string
	Role        string
	Status      string
}

type UpdateProfile struct {
	Name        string
	Email       string
	DateOfBirth string
	Gender      string
	Avatar      string
	Phone       string
}

type UpdatePassword struct {
	Password string
}

type PatientCredential struct {
	Name   string
	Email  string
	Access map[string]any
}

type AvatarPhoto struct {
	Avatar multipart.File
}

type PatientHandlerInterface interface {
	GetPatients() echo.HandlerFunc
	CreatePatient() echo.HandlerFunc
	UpdatePatient() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
	LoginPatient() echo.HandlerFunc
}

type PatientServiceInterface interface {
	GetPatients() ([]Patientdetail, error)
	GetPatient(id int) (Patientdetail, error)
	CreatePatient(newData Patiententity) (*Patiententity, error)
	PhotoUpload(newData AvatarPhoto) (string, error)
	LoginPatient(email string, password string) (*PatientCredential, error)
	UpdatePatient(id int, newData UpdateProfile) (bool, error)
	UpdatePassword(id int, newData UpdatePassword) (bool, error)
}

type PatientDataInterface interface {
	GetAll() ([]Patientdetail, error)
	GetByID(id int) (Patientdetail, error)
	Insert(newData Patiententity) (*Patiententity, error)
	LoginPatient(email string, password string) (*Patiententity, error)
	Update(id int, newData UpdateProfile) (bool, error)
	UpdatePassword(id int, newData UpdatePassword) (bool, error)
}
