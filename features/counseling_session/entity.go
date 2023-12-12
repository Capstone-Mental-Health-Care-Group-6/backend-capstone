package CounselingSession

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CounselingSession struct {
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
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type CounselingSessionInfo struct {
	TransactionID   uint              `json:"transaction_id"`
	TransactionInfo []TransactionInfo `json:"transaction_info"`
	Date            time.Time         `json:"date"`
	Time            time.Time         `json:"time"`
	Duration        uint              `json:"duration"`
	Status          string            `json:"status"`
}

type TransactionInfo struct {
	TopicName     string `json:"topic_name"`
	PatientName   string `json:"patient_name"`
	PatientAvatar string `json:"patient_avatar"`
	DoctorName    string `json:"doctor_name"`
	MethodName    string `json:"method_name"`
	DurationName  string `json:"duration_name"`
	CounselingID  uint   `json:"counseling_id"`

	UserID     uint   `json:"user_id"`
	MidtransID string `json:"transaction_id"`

	CounselingType string `json:"counseling_type"`

	PriceMethod     uint `json:"price_method"`
	PriceDuration   uint `json:"price_duration"`
	PriceCounseling uint `json:"price_counseling"`
	PriceResult     uint `json:"price_result"`

	PaymentProof  string    `json:"payment_proof"`
	PaymentStatus uint      `json:"payment_status"`
	PaymentType   string    `json:"payment_type"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
	// Ratings       []DoctorRating `json:"ratings gorm:"foreignkey:DoctorID"`
}

type CounselingSessionHandlerInterface interface {
	GetAllCounseling() echo.HandlerFunc
	CreateCounseling() echo.HandlerFunc
	GetCounseling() echo.HandlerFunc
	GetCounselingByUserID() echo.HandlerFunc
	UpdateCounseling() echo.HandlerFunc
	DeleteCounseling() echo.HandlerFunc
}

type CounselingSessionServiceInterface interface {
	GetAllCounseling() ([]CounselingSession, error)
	CreateCounseling(input CounselingSession) (*CounselingSession, error)
	GetCounseling(id int) (*CounselingSession, error)
	GetAllCounselingByUserID(userID int) ([]CounselingSession, error)
	UpdateCounseling(id int, input CounselingSession) (bool, error)
	DeleteCounseling(id int) (bool, error)
}

type CounselingSessionDataInterface interface {
	GetAll() ([]CounselingSession, error)
	Create(input CounselingSession) (*CounselingSession, error)
	GetAllCounselingByUserID(userID int) ([]CounselingSession, error)
	GetById(id int) (*CounselingSession, error)
	Update(id int, newData CounselingSession) (bool, error)
	Delete(id int) (bool, error)
}
