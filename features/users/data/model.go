package data

import (
	articles "FinalProject/features/articles/data"
	patient "FinalProject/features/patients/data"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name           string             `gorm:"column:name;type:varchar(255)"`
	Email          string             `gorm:"column:email;unique;type:varchar(255)"`
	Password       string             `gorm:"column:password;type:varchar(255)"`
	Role           string             `gorm:"type:enum('Admin','Patient','Doctor');column:role"`
	Status         string             `gorm:"type:enum('Active','Inactive','Suspend');column:status"`
	TokenResetPass string             `gorm:"column:token_reset_pass;type:varchar(255)"`
	Articles       []articles.Article `gorm:"foreignKey:UserID"`
	Patient        patient.Patient    `gorm:"foreignKey:UserID;references:ID"`
}
