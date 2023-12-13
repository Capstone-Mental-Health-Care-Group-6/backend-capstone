package counselingdurations

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CounselingDuration struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	AdditionalPrice int       `json:"additional_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CounselingDurationInfo struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	AdditionalPrice int    `json:"additional_price"`
}

type CounselingDurationHandlerInterface interface {
	GetCounselingDurations() echo.HandlerFunc
	GetCounselingDuration() echo.HandlerFunc
}

type CounselingDurationServiceInterface interface {
	GetAll() ([]CounselingDurationInfo, error)
	GetByID(id int) ([]CounselingDurationInfo, error)
}

type CounselingDurationDataInterface interface {
	GetAll() ([]CounselingDurationInfo, error)
	GetByID(id int) ([]CounselingDurationInfo, error)
}
