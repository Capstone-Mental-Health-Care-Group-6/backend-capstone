package data

import "gorm.io/gorm"

type BundleCounseling struct {
	*gorm.Model
	Name         string `gorm:"column:name"`
	Sessions     uint   `gorm:"column:sessions"`
	Type         string `gorm:"column:type;type:enum('PREMIUM','INSTAN')"`
	Price        uint   `gorm:"column:price"`
	Description  string `gorm:"column:description"`
	ActivePriode string `gorm:"column:active_priode"`
}

func (BundleCounseling) TableName() string {
	return "bundle_counseling"
}
