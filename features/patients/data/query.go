package data

import (
	//"net/url"
	"FinalProject/features/patients"
	"FinalProject/helper"
	//mysql "FinalProject/utils/database/migration/mysql"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PatientData struct {
	db *gorm.DB
}

func New(db *gorm.DB) patients.PatientDataInterface {
	return &PatientData{
		db: db,
	}
}

func (pdata *PatientData) GetAll() ([]patients.Patientdetail, error) {
	//db := pdata.db.Table("patient_accounts")
	//db = helper.QueryFiltering(db, query)
	//db = helper.QuerySorting(db, query)
	//db = helper.QueryPagination(db, query)
	//data := make([]patients.Patientdetail, 0)
	//if err := db.where("deleted_at is null").Find(&data).Error; err != nil {
	//	logrus.Error("[PatientData.GetAll] Error : ", err.Error())
	//	return nil
	//}

	var listPatient = []patients.Patientdetail{}
	var qry = pdata.db.Table("patient_accounts").Select("patient_accounts.*").
		Where("patient_accounts.deleted_at is null").
		Scan(&listPatient)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listPatient, nil
}

func (pdata *PatientData) GetByID(id int) (patients.Patientdetail, error) {
	var listPatient patients.Patientdetail
	var qry = pdata.db.Table("patient_accounts").Select("patient_accounts.*").
		Where("patient_accounts.id = ?", id).
		Where("patient_accounts.deleted_at is null").
		Scan(&listPatient)

	if err := qry.Error; err != nil {
		return listPatient, err
	}
	return listPatient, nil
}

func (pdata *PatientData) Insert(newData patients.Patiententity) (*patients.Patiententity, error) {

	var dbData = new(PatientAccount)
	dbData.Name = newData.Name
	dbData.Email = newData.Email
	dbData.DateOfBirth = newData.DateOfBirth
	dbData.Gender = newData.Gender
	dbData.Avatar = newData.Avatar
	dbData.Phone = newData.Phone
	dbData.Role = "Patient"
	dbData.Status = "Active"
	hashPassword, err := helper.HashPassword(newData.Password)
	if err != nil {
		logrus.Info("Hash Password Error, ", err.Error())
	}
	dbData.Password = hashPassword

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (pdata *PatientData) LoginPatient(email, password string) (*patients.Patiententity, error) {
	var dbData = new(PatientAccount)
	dbData.Email = email

	if err := pdata.db.Where("email = ?", dbData.Email).First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), passwordBytes)
	if err != nil {
		logrus.Info("Incorrect Password")
		return nil, err
	}

	var result = new(patients.Patiententity)
	result.ID = dbData.ID
	result.Email = dbData.Email
	result.Name = dbData.Name
	result.Role = dbData.Role
	result.Status = dbData.Status

	return result, nil
}

func (pdata *PatientData) Update(id int, newData patients.UpdateProfile) (bool, error) {
	var qry = pdata.db.Table("patient_accounts").Where("id = ?", id).Updates(PatientAccount{
		Name:        newData.Name,
		Email:       newData.Email,
		DateOfBirth: newData.DateOfBirth,
		Gender:      newData.Gender,
		Avatar:      newData.Avatar,
		Phone:       newData.Phone,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, nil
	}

	return true, nil
}

func (pdata *PatientData) UpdatePassword(id int, newData patients.UpdatePassword) (bool, error) {
	var qry = pdata.db.Table("patient_accounts").Where("id = ?", id).Updates(PatientAccount{
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
