package seeds

import (
	counselingduration "FinalProject/features/counseling_durations"
	"time"

	"gorm.io/gorm"
)

func CreateCounselingDuration(db *gorm.DB, name string, price int) error {
	return db.Create(&counselingduration.CounselingDuration{Name: name, AdditionalPrice: price, CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error
}
