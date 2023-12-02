package data

import (
	"FinalProject/features/users"
	"FinalProject/helper"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}

func (ud *UserData) Register(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Name = newData.Name
	dbData.Email = newData.Email
	dbData.Role = newData.Role
	dbData.Status = "Active"
	hashPassword, err := helper.HashPassword(newData.Password)
	if err != nil {
		logrus.Info("Hash Password Error, ", err.Error())
	}
	dbData.Password = hashPassword

	if err := ud.db.Create(dbData).Error; err != nil {
		return nil, err
	}
	return &newData, nil
}

func (ud *UserData) Login(email, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email

	if err := ud.db.Where("email = ?", dbData.Email).First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), passwordBytes)
	if err != nil {
		logrus.Info("Incorrect Password")
		return nil, errors.New("Incorrect Password")
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Email = dbData.Email
	result.Name = dbData.Name
	result.Role = dbData.Role
	result.Status = dbData.Status

	return result, nil
}

func (ud *UserData) GetByEmail(email string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email

	if err := ud.db.Where("email = ?", dbData.Email).First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Email = dbData.Email
	result.Name = dbData.Name
	result.Role = dbData.Role
	result.Status = dbData.Status

	return result, nil
}

func (ud *UserData) GetUsers(status, name string) ([]users.User, error) {
	var listUser = []users.User{}

	var qry = ud.db.Table("users").Where("role = ?", "Patient")

	if status != "" {
		qry.Where("status = ?", status)
	}

	if name != "" {
		qry.Where("name LIKE ?", "%"+name+"%")
	}

	if err := qry.Scan(&listUser).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	return listUser, nil
}

func (ud *UserData) UserDashboard() (users.UserDashboard, error) {
	var dashboardUser users.UserDashboard

	tUser, tUserBaru, tUserActive, tUserInactive := ud.getTotalUser()

	dashboardUser.TotalUser = tUser
	dashboardUser.TotalUserBaru = tUserBaru
	dashboardUser.TotalUserActive = tUserActive
	dashboardUser.TotalUserInactive = tUserInactive

	return dashboardUser, nil
}

func (ud *UserData) getTotalUser() (int, int, int, int) {
	var totalUser int64
	var totalUserBaru int64
	var totalUserActive int64
	var totalUserInactive int64

	var now = time.Now()
	var before = now.AddDate(0, 0, -30)

	var _ = ud.db.Table("users").Where("role = ?", "Patient").Count(&totalUser)
	var _ = ud.db.Table("users").Where("role = ?", "Patient").Where("created_at BETWEEN ? and ?", before, now).Count(&totalUserBaru)
	var _ = ud.db.Table("users").Where("role = ?", "Patient").Where("status = ?", "Active").Count(&totalUserActive)
	var _ = ud.db.Table("users").Where("role = ?", "Patient").Where("status = ?", "Inactive").Count(&totalUserInactive)

	totalUserInt := int(totalUser)
	totalUserBaruInt := int(totalUserBaru)
	totalUserActiveInt := int(totalUserActive)
	totalUserInactiveInt := int(totalUserInactive)

	return totalUserInt, totalUserBaruInt, totalUserActiveInt, totalUserInactiveInt
}
