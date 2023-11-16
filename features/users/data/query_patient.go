package data

import (
	"FinalProject/features/users"
	"gorm.io/gorm"
)

type PatientData struct {
	db *gorm.DB
}

func NewPatient(db *gorm.DB) users.PatientDataInterface {
	return &PatientData{
		db: db,
	}
}

func (pdata *PatientData) GetAll() ([]users.Patiententity, error) {
	var listPatient = []users.Patiententity{}
	var qry = pdata.db.Table("patients").Select("patients.*").
		Where("patients.deleted_at is null").
		Scan(&listPatient)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listPatient, nil
}

func (pdata *PatientData) GetByID(id int) ([]users.Patiententity, error) {
	var listPatient = []users.Patiententity{}
	var qry = pdata.db.Table("patients").Select("patients.*").
		Where("patients.id = ?", id).
		Where("patients.deleted_at is null").
		Scan(&listPatient)

	if err := qry.Error; err != nil {
		return nil, err
	}
	return listPatient, nil
}

func (pdata *PatientData) Insert(newData users.Patiententity) (*users.Patiententity, error) {
	var dbData = new(Patient)
	dbData.Name = newData.Name
	dbData.UserID = newData.UserID
	dbData.DateOfBirth = newData.DateOfBirth
	dbData.PlaceOfBirth = newData.PlaceOfBirth
	dbData.Gender = newData.Gender
	dbData.MarriageStatus = newData.MarriageStatus
	dbData.Avatar = newData.Avatar
	dbData.Address = newData.Address

	//handling error for duplicate user id won't fix
	if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
		return nil, err
	}

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}
