package data

import (
	"FinalProject/features/doctor"
	"fmt"

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

func (pdata *DoctorData) GetAll() ([]doctor.DoctorInfo, error) {
	var listDoctor = []doctor.DoctorInfo{}
	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id, doctors_workadays.workday_id AS workday_id, doctors_workadays.start_time AS start_time, doctors_workadays.end_time AS end_time").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors.id = doctors_expertise_relation.doctor_id").
		Joins("LEFT JOIN doctors_workadays ON doctors.id = doctors_workadays.doctor_id").
		Where("doctors.deleted_at IS NULL").
		Scan(&listDoctor)

	fmt.Println("List doctor:", listDoctor)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listDoctor, nil
}

func (pdata *DoctorData) GetByID(id int) ([]doctor.DoctorInfo, error) {
	var listDoctor []doctor.DoctorInfo

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id, doctors_workadays.workday_id AS workday_id, doctors_workadays.start_time AS start_time, doctors_workadays.end_time AS end_time").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors.id = doctors_expertise_relation.doctor_id").
		Joins("LEFT JOIN doctors_workadays ON doctors.id = doctors_workadays.doctor_id").
		Where("doctors.id = ?", id).
		Where("doctors.deleted_at IS NULL").
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
	dbData.DoctorExperienced = newData.DoctorExperienced
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

	// var dbDataExpertise = new(DoctorExpertiseRelation)
	// dbDataExpertise.DoctorID = // THE DOCTOR ID ABOVE
	// dbDataExpertise.ExpertiseID = new

	//handling error for duplicate user id won't fix
	if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
		return nil, err
	}

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	newData.ID = dbData.ID

	return &newData, nil
}

func (pdata *DoctorData) InsertExpertise(newData doctor.DoctorExpertiseRelation) (*doctor.DoctorExpertiseRelation, error) {

	var dbData = new(DoctorExpertiseRelation)
	dbData.DoctorID = newData.DoctorID
	dbData.ExpertiseID = newData.ExpertiseID

	//handling error for duplicate user id won't fix
	// if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
	// 	return nil, err
	// }

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (pdata *DoctorData) InsertWorkadays(newData doctor.DoctorWorkadays) (*doctor.DoctorWorkadays, error) {

	var dbData = new(DoctorWorkadays)
	dbData.DoctorID = newData.DoctorID
	dbData.WorkdayID = newData.WorkdayID
	dbData.StartTime = newData.StartTime
	dbData.EndTime = newData.EndTime
	//handling error for duplicate user id won't fix
	// if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
	// 	return nil, err
	// }

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (pdata *DoctorData) InsertExperience(newData doctor.DoctorExperience) (*doctor.DoctorExperience, error) {

	var dbData = new(DoctorExperience)
	dbData.DoctorID = newData.DoctorID
	dbData.DoctorCompany = newData.DoctorCompany
	dbData.DoctorTitle = newData.DoctorTitle
	dbData.DoctorExperienceDescription = newData.DoctorExperienceDescription
	dbData.DoctorStartDate = newData.DoctorStartDate
	dbData.DoctorEndDate = newData.DoctorEndDate
	dbData.DoctorIsNow = newData.DoctorIsNow
	//handling error for duplicate user id won't fix
	// if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
	// 	return nil, err
	// }

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (pdata *DoctorData) InsertEducation(newData doctor.DoctorEducation) (*doctor.DoctorEducation, error) {

	var dbData = new(DoctorEducation)
	dbData.DoctorID = newData.DoctorID
	dbData.DoctorUniversity = newData.DoctorUniversity
	dbData.DoctorStudyProgram = newData.DoctorStudyProgram
	dbData.DoctorGraduateYear = newData.DoctorGraduateYear
	//handling error for duplicate user id won't fix
	// if err := pdata.db.Where("user_id = ?", dbData.UserID).Find(dbData).Error; err != nil {
	// 	return nil, err
	// }

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}
