package data

import (
	"FinalProject/features/doctor"

	"gorm.io/gorm"
)

type DoctorData struct {
	db *gorm.DB
}

func NewDoctor(db *gorm.DB) doctor.DoctorDataInterface {
	return &DoctorData{
		db: db,
	}
}

func (pdata *DoctorData) GetAll() ([]doctor.Doctor, error) {
	var listDoctor = []doctor.Doctor{}
	var qry = pdata.db.Table("doctors").Select("doctors.*").
		Where("doctors.deleted_at is null").
		Scan(&listDoctor)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listDoctor, nil
}

func (pdata *DoctorData) GetByID(id int) ([]doctor.Doctor, error) {
	var listDoctor = []doctor.Doctor{}
	var qry = pdata.db.Table("doctors").Select("doctors.*").
		Where("doctors.id = ?", id).
		Where("doctors.deleted_at is null").
		Scan(&listDoctor)

	if err := qry.Error; err != nil {
		return nil, err
	}
	return listDoctor, nil
}

func (pdata *DoctorData) Insert(newData doctor.Doctor) (*doctor.Doctor, error) {

	var dbData = new(Doctor)
	dbData.DoctorName = newData.DoctorName
	dbData.UserID = newData.UserID
	dbData.DoctorExperience = newData.DoctorExperience
	dbData.DoctorDescription = newData.DoctorDescription
	dbData.DoctorAvatar = newData.DoctorAvatar
	dbData.DoctorOfficeName = newData.DoctorOfficeName
	dbData.DoctorOfficeAddress = newData.DoctorOfficeAddress
	dbData.DoctorOfficeCity = newData.DoctorOfficeCity
	dbData.DoctorMeetLink = newData.DoctorMeetLink
	dbData.DoctorSIPP = newData.DoctorSIPP
	dbData.DoctorSIPPFile = newData.DoctorSIPPFile
	dbData.DoctorSTR = newData.DoctorSTR
	dbData.DoctorSTRFile = newData.DoctorSTRFile
	dbData.DoctorCV = newData.DoctorCV
	dbData.DoctorIjazah = newData.DoctorIjazah
	dbData.DoctorBalance = newData.DoctorBalance
	dbData.DoctorStatus = newData.DoctorStatus

	//handling error for duplicate user id won't fix
	if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
		return nil, err
	}

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}
