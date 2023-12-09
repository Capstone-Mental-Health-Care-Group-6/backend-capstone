package data

import (
	dataArticle "FinalProject/features/articles/data"
	dataDoctor "FinalProject/features/doctor/data"
	"time"

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

type UserResetPass struct {
	*gorm.Model
	Email     string    `json:"email" gorm:"email"`
	Code      string    `json:"code"  gorm:"code"`
	ExpiresAt time.Time `json:"expiresat" gorm:"expiresat"`
}

func (User) TableName() string {
	return "users"
}

func (UserResetPass) TableName() string {
	return "user_reset_pass"
}
