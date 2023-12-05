package data

import (
	"time"

	"gorm.io/gorm"
)

type CounselingSession struct {
	*gorm.Model
	TransactionID uint      `gorm:"column:transaction_id"`
	Date          time.Time `gorm:"column:date"`
	Time          time.Time `gorm:"column:time"`
	Duration      uint      `gorm:"column:duration"`
	Status        string    `gorm:"column:status;type:enum('not_finished','finished')"`
}

func (CounselingSession) TableName() string {
	return "counseling_session"
}
