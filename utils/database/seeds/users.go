package seeds

import (
	"FinalProject/features/users"
	"FinalProject/helper"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name, email, password, role, status string) error {
	hashPassword, _ := helper.HashPassword(password)

	return db.Create(&users.User{Name: name, Email: email, Password: hashPassword, Role: role, Status: status}).Error
}
