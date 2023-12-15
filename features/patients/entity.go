package patients

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
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
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Avatar      string `json:"avatar"`
	Phone       string `json:"phone_number"`
	Role        string `json:"role"`
	Status      string `json:"status"`
}

type UpdateProfile struct {
	Name        string `json:"name"`
	Email       string
	DateOfBirth string
	Gender      string
	Avatar      string
	Phone       string
}

type UpdatePassword struct {
	Password string
}

type UpdateStatus struct {
	Status string
}

type PatientCredential struct {
	Name   string
	Email  string
	Access map[string]any
}

type AvatarPhoto struct {
	Avatar multipart.File
}

type PatientDashboard struct {
	TotalUser         int `json:"total_user"`
	TotalUserBaru     int `json:"total_user_baru"`
	TotalUserActive   int `json:"total_user_active"`
	TotalUserInactive int `json:"total_user_inactive"`
}

type PatientHandlerInterface interface {
	GetPatients() echo.HandlerFunc
	GetPatient() echo.HandlerFunc
	CreatePatient() echo.HandlerFunc
	UpdatePatient() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
	UpdateStatus() echo.HandlerFunc
	LoginPatient() echo.HandlerFunc
	PatientDashboard() echo.HandlerFunc
	InactivateAccount() echo.HandlerFunc
}

type PatientServiceInterface interface {
	GetPatients(status, name string) ([]Patientdetail, error)
	GetPatient(id int) (Patientdetail, error)
	CreatePatient(newData Patiententity) (*Patiententity, error)
	PhotoUpload(newData AvatarPhoto) (string, error)
	LoginPatient(email string, password string) (*PatientCredential, error)
	UpdatePatient(id int, newData UpdateProfile) (bool, error)
	UpdatePassword(id int, newData UpdatePassword) (bool, error)
	PatientDashboard() (PatientDashboard, error)
	UpdateStatus(id int, newData UpdateStatus) (bool, error)
	InactivateAccount(id int) (bool, error)
}

type PatientDataInterface interface {
	GetAll(status, name string) ([]Patientdetail, error)
	GetByID(id int) (Patientdetail, error)
	Insert(newData Patiententity) (*Patiententity, error)
	LoginPatient(email string, password string) (*Patiententity, error)
	Update(id int, newData UpdateProfile) (bool, error)
	UpdatePassword(id int, newData UpdatePassword) (bool, error)
	PatientDashboard() (PatientDashboard, error)
	UpdateStatus(id int, newData UpdateStatus) (bool, error)
	InactivateAccount(id int) (bool, error)
}
