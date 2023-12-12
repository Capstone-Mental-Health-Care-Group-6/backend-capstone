package data

import (
	"time"

	"gorm.io/gorm"
)

type CounselingSession struct {
	*gorm.Model
	TransactionID   uint      `gorm:"column:transaction_id"`
	UserID          uint      `gorm:"column:user_id"`
	DoctorAvatar    string    `gorm:"doctor_avatar"`
	DoctorName      string    `gorm:"doctor_name"`
	DoctorExpertise uint      `gorm:"doctor_expertise"`
	DoctorMeetLink  string    `gorm:"doctor_meet_link"`
	Date            time.Time `gorm:"column:date"`
	Time            time.Time `gorm:"column:time"`
	Duration        uint      `gorm:"column:duration"`
	Status          string    `gorm:"column:status;type:enum('process','done')"`
}

func (CounselingSession) TableName() string {
	return "counseling_session"
}
