package data

import "gorm.io/gorm"

type CounselingDuration struct {
	*gorm.Model
	Name            string `gorm:"column:name;type:varchar(255)"`
	AdditionalPrice string `gorm:"column:additional_price;type:int(11)"`
}
