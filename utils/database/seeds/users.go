package seeds

import (
	"FinalProject/features/users"
	"FinalProject/helper/enkrip"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name, email, password, role, status string) error {
	hash := enkrip.New()
	hashPassword, _ := hash.HashPassword(password)

	return db.Create(&users.User{Name: name, Email: email, Password: hashPassword, Role: role, Status: status}).Error
}
