package seeds

import (
	counselingtopics "FinalProject/features/counseling_topics"
	"time"

	"gorm.io/gorm"
)

func CreateCounselingTopic(db *gorm.DB, name string) error {
	return db.Create(&counselingtopics.CounselingTopic{Name: name, CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error
}
