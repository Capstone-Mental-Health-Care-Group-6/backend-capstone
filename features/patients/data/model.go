package data

import (
	"gorm.io/gorm"
)

type PatientAccount struct {
	*gorm.Model
	Name        string `gorm:"column:name;type:varchar(255)"`
	Email       string `gorm:"column:email;unique;type:varchar(255)"`
	Password    string `gorm:"column:password;type:varchar(255)"`
	DateOfBirth string `gorm:"column:date_of_birth;type:date"`
	Gender      string `gorm:"column:gender;type:enum('Laki-laki','Perempuan')"`
	Avatar      string `gorm:"column:avatar;type:varchar(255)"`
	Phone       string `gorm:"column:phone;type:varchar(255)"`
	Role        string `gorm:"column:role;type:varchar(255)"`
	Status      string `gorm:"column:status;type:enum('Active','Inactive')"`
}

func (PatientAccount) TableName() string {
	return "patient_accounts"
}
