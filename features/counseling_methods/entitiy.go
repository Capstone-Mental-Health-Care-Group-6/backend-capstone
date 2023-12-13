package counselingmethods

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CounselingMethod struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	AdditionalPrice int       `json:"additional_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CounselingMethodInfo struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	AdditionalPrice int    `json:"additional_price"`
}

type CounselingMethodHandlerInterface interface {
	GetCounselingMethods() echo.HandlerFunc
	GetCounselingMethod() echo.HandlerFunc
}

type CounselingMethodServiceInterface interface {
	GetAll() ([]CounselingMethodInfo, error)
	GetByID(id int) ([]CounselingMethodInfo, error)
}

type CounselingMethodDataInterface interface {
	GetAll() ([]CounselingMethodInfo, error)
	GetByID(id int) ([]CounselingMethodInfo, error)
}
