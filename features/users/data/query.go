package data

import (
	"FinalProject/features/users"
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
	dbData.Status = newData.Status
	dbData.Password = newData.Password

	if err := ud.db.Create(dbData).Error; err != nil {
		return nil, err
	}
	return &newData, nil
}

func (ud *UserData) Login(email, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email

	var qry = ud.db.Where("email = ? AND status = ?", dbData.Email, "Active").First(dbData)

	var dataCount int64
	qry.Count(&dataCount)
	if dataCount == 0 {
		return nil, errors.New("Not Found")
	}

	if err := qry.Error; err != nil {
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

func (ud *UserData) GetByID(id int) (users.User, error) {
	var listUser users.User
	var qry = ud.db.Table("users").Select("users.*").
		Where("users.id = ?", id).
		Where("users.deleted_at is null").
		Scan(&listUser)

	if err := qry.Error; err != nil {
		return listUser, err
	}
	return listUser, nil
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

func (ud *UserData) InsertCode(email string, code string) error {
	var newData = new(UserResetPass)
	newData.Email = email
	newData.Code = code
	newData.ExpiresAt = time.Now().Add(10 * time.Minute)

	_, err := ud.GetByCode(code)
	if err != nil {
		ud.DeleteCode(code)
	}

	if err := ud.db.Table("user_reset_pass").Create(newData).Error; err != nil {
		return err
	}

	return nil
}

func (ud *UserData) DeleteCode(code string) error {
	var deleteData = new(UserResetPass)

	if err := ud.db.Table("user_reset_pass").Where("code = ?", code).Delete(deleteData).Error; err != nil {
		return err
	}

	return nil
}

func (ud *UserData) GetByCode(code string) (*users.UserResetPass, error) {
	var dbData = new(UserResetPass)
	dbData.Code = code

	if err := ud.db.Table("user_reset_pass").Where("code = ?", dbData.Code).Where("deleted_at IS NULL").First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	var result = new(users.UserResetPass)
	result.Email = dbData.Email
	result.Code = dbData.Code
	result.ExpiresAt = dbData.ExpiresAt

	return result, nil
}

func (ud *UserData) ResetPassword(code, email string, password string) error {
	if err := ud.db.Table("users").Where("email = ?", email).Update("password", password).Error; err != nil {
		return err
	}

	checkData, _ := ud.GetByCode(code)
	if checkData.Code != "" {
		ud.DeleteCode(code)
	}

	return nil
}

func (ud *UserData) UpdateProfile(id int, newData users.UpdateProfile) (bool, error) {
	var qry = ud.db.Table("users").Where("id = ?", id).Updates(User{
		Name:     newData.Name,
		Email:    newData.Email,
		Password: newData.Password,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, nil
	}

	return true, nil
}
