package seeds

import (
	counselingmethod "FinalProject/features/counseling_methods"
	"time"

	"gorm.io/gorm"
)

func CreateCounselingMethod(db *gorm.DB, name string, price int) error {
	return db.Create(&counselingmethod.CounselingMethod{Name: name, AdditionalPrice: price, CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error
}
