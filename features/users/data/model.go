package data

import (
	dataArticle "FinalProject/features/articles/data"
	dataDoctor "FinalProject/features/doctor/data"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name           string                `gorm:"column:name;type:varchar(255)"`
	Email          string                `gorm:"column:email;unique;type:varchar(255)"`
	Password       string                `gorm:"column:password;type:varchar(255)"`
	Role           string                `gorm:"type:enum('Admin','Patient','Doctor');column:role"`
	Status         string                `gorm:"type:enum('Active','Inactive','Suspend');column:status"`
	TokenResetPass string                `gorm:"column:token_reset_pass;type:varchar(255)"`
	Articles       []dataArticle.Article `gorm:"foreignKey:UserID"`
	Doctors        []dataDoctor.Doctor   `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
