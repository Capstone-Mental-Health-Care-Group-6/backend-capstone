package data

import (
	"gorm.io/gorm"
)

type Patient struct {
	*gorm.Model
	Name           string `gorm:"column:name;type:varchar(255)"`
	UserID         uint   `gorm:"column:user_id"`
	DateOfBirth    string `gorm:"column:date_of_birth;type:date"`
	PlaceOfBirth   string `gorm:"column:place_of_birth;type:varchar(255)"`
	Gender         string `gorm:"column:gender;type:enum('Male','Female')"`
	MarriageStatus string `gorm:"column:marriage_status;type:enum('Single','Married','Divorced','Widowed')"`
	Avatar         string `gorm:"column:avatar;type:varchar(255)"`
	Address        string `gorm:"column:address;type:varchar(255)"`
}
