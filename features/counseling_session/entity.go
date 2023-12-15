package CounselingSession

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CounselingSession struct {
	ID              uint      `json:"id"`
	TransactionID   uint      `json:"transaction_id"`
	UserID          uint      `json:"user_id"`
	DoctorAvatar    string    `json:"doctor_avatar"`
	DoctorName      string    `json:"doctor_name"`
	DoctorExpertise uint      `json:"doctor_expertise"`
	DoctorMeetLink  string    `json:"doctor_meet_link"`
	Date            time.Time `json:"date"`
	Time            time.Time `json:"time"`
	Duration        uint      `json:"duration"`
	Status          string    `json:"status"`
	Alasan          string    `json:"alasan"`
	DetailAlasan    string    `json:"detail_alasan"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`

	CounselingType   string `json:"counseling_type"`
	CounselingMethod string `json:"counseling_method"`
	CounselingTopic  string `json:"counseling_topic"`
}

type StatusUpdate struct {
	Alasan string `json:"alasan"`
}

type CounselingSessionHandlerInterface interface {
	GetAllCounseling() echo.HandlerFunc
	CreateCounseling() echo.HandlerFunc
	GetCounseling() echo.HandlerFunc
	GetCounselingByUserID() echo.HandlerFunc
	UpdateCounseling() echo.HandlerFunc
	DeleteCounseling() echo.HandlerFunc
	RejectPatient() echo.HandlerFunc
	ApprovePatient() echo.HandlerFunc
}

type CounselingSessionServiceInterface interface {
	GetAllCounseling() ([]CounselingSession, error)
	CreateCounseling(input CounselingSession) (*CounselingSession, error)
	GetCounseling(id int) (*CounselingSession, error)
	GetAllCounselingByUserID(userID int) ([]CounselingSession, error)
	UpdateCounseling(id int, input CounselingSession) (bool, error)
	DeleteCounseling(id int) (bool, error)
	ApprovePatient(id, doctorID int) (bool, error)
	RejectPatient(id, doctorID int, newData StatusUpdate) (bool, error)
}

type CounselingSessionDataInterface interface {
	GetAll() ([]CounselingSession, error)
	Create(input CounselingSession) (*CounselingSession, error)
	GetAllCounselingByUserID(userID int) ([]CounselingSession, error)
	GetById(id int) (*CounselingSession, error)
	Update(id int, newData CounselingSession) (bool, error)
	Delete(id int) (bool, error)
	ApprovePatient(id int) (bool, error)
	RejectPatient(id int, newData StatusUpdate) (bool, error)
	CheckPatient(id, doctorID int) error
}
