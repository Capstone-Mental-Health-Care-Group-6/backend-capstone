package data

import (
	"FinalProject/features/patients"
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

func (pdata *PatientData) GetAll() ([]patients.Patiententity, error) {
	var listPatient = []patients.Patiententity{}
	var qry = pdata.db.Table("patients").Select("patients.*").
		Where("patients.deleted_at is null").
		Scan(&listPatient)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listPatient, nil
}

func (pdata *PatientData) GetByID(id int) ([]patients.Patiententity, error) {
	var listPatient = []patients.Patiententity{}
	var qry = pdata.db.Table("patients").Select("patients.*").
		Where("patients.id = ?", id).
		Where("patients.deleted_at is null").
		Scan(&listPatient)

	if err := qry.Error; err != nil {
		return nil, err
	}
	return listPatient, nil
}

func (pdata *PatientData) Insert(newData patients.Patiententity) (*patients.Patiententity, error) {
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
