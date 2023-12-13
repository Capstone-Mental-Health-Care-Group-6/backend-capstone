package counselingtopics

import (
	"time"

	"github.com/labstack/echo/v4"
)

type CounselingTopic struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CounselingTopicInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CounselingTopicHandlerInterface interface {
	GetCounselingTopics() echo.HandlerFunc
	GetCounselingTopic() echo.HandlerFunc
}

type CounselingTopicServiceInterface interface {
	GetAll() ([]CounselingTopicInfo, error)
	GetByID(id int) ([]CounselingTopicInfo, error)
}

type CounselingTopicDataInterface interface {
	GetAll() ([]CounselingTopicInfo, error)
	GetByID(id int) ([]CounselingTopicInfo, error)
}
