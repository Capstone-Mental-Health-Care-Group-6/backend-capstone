package CounselingSession

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CounselingSession struct {
	TransactionID uint      `json:"transaction_id"`
	Date          time.Time `json:"date"`
	Time          time.Time `json:"time"`
	Duration      uint      `json:"duration"`
	Status        string    `json:"status"`
}

type CounselingSessionHandlerInterface interface {
	GetAllCounseling() echo.HandlerFunc
	CreateCounseling() echo.HandlerFunc
	GetCounseling() echo.HandlerFunc
	UpdateCounseling() echo.HandlerFunc
	DeleteCounseling() echo.HandlerFunc
}

type CounselingSessionServiceInterface interface {
	GetAllCounseling() ([]CounselingSession, error)
	CreateCounseling(input CounselingSession) (*CounselingSession, error)
	GetCounseling(id int) (*CounselingSession, error)
	UpdateCounseling(id int, input CounselingSession) (bool, error)
	DeleteCounseling(id int) (bool, error)
}

type CounselingSessionDataInterface interface {
	GetAll() ([]CounselingSession, error)
	Create(input CounselingSession) (*CounselingSession, error)
	GetById(id int) (*CounselingSession, error)
	Update(id int, newData CounselingSession) (bool, error)
	Delete(id int) (bool, error)
}
