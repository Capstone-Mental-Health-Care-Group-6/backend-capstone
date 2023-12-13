package data

import (
	"FinalProject/features/doctor"
	"FinalProject/features/users"
	"fmt"
	"github.com/sirupsen/logrus"
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

// SEARCH QUERY \\

func (pdata *DoctorData) SearchDoctor(name string) ([]doctor.DoctorAll, error) {
	var doctors []doctor.DoctorAll

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors.id = doctors_expertise_relation.doctor_id").
		Where("doctors.deleted_at IS NULL AND doctors.doctor_name LIKE ?", "%"+name+"%").
		Scan(&doctors)

	if err := qry.Error; err != nil {
		return nil, err
	}

	for i, doctor := range doctors {
		// Retrieve experience, education, and workday for each doctor
		experience, err := pdata.GetByIDExperience(int(doctor.ID))
		if err != nil {
			return nil, err
		}
		education, err := pdata.GetByIDEducation(int(doctor.ID))
		if err != nil {
			return nil, err
		}
		workday, err := pdata.GetByIDWorkadays(int(doctor.ID))
		if err != nil {
			return nil, err
		}

		// Assign the retrieved data to the corresponding fields in the Doctor struct
		doctors[i].DoctorExperience = experience
		doctors[i].DoctorEducation = education
		doctors[i].DoctorWorkday = workday
	}

	return doctors, nil
}

// FIND EMAIL QUERY \\

func (pdata *DoctorData) FindEmail(userID uint) (*string, error) {

	var dbData = new(users.User)

	if err := pdata.db.Table("users").Where("id = ?", userID).First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.Email = dbData.Email

	return &result.Email, nil
}

// GET ALL AND BY ID QUERY \\

func (pdata *DoctorData) GetAll() ([]doctor.DoctorAll, error) {
	var doctors []doctor.DoctorAll

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors.id = doctors_expertise_relation.doctor_id").
		Where("doctors.deleted_at IS NULL").
		Scan(&doctors)

	if err := qry.Error; err != nil {
		return nil, err
	}

	for i, doctor := range doctors {
		// Retrieve experience, education, and workday for each doctor
		experience, err := pdata.GetByIDExperience(int(doctor.ID))
		if err != nil {
			return nil, err
		}
		education, err := pdata.GetByIDEducation(int(doctor.ID))
		if err != nil {
			return nil, err
		}
		workday, err := pdata.GetByIDWorkadays(int(doctor.ID))
		if err != nil {
			return nil, err
		}

		fmt.Println("Data", doctor.ID)

		// Assign the retrieved data to the corresponding fields in the Doctor struct
		doctors[i].DoctorExperience = experience
		doctors[i].DoctorEducation = education
		doctors[i].DoctorWorkday = workday
	}

	return doctors, nil
}

func (pdata *DoctorData) GetByID(id int) (*doctor.DoctorAll, error) {
	var doctor doctor.DoctorAll

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors.id = doctors_expertise_relation.doctor_id").
		Where("doctors.id = ?", id).
		Where("doctors.deleted_at IS NULL").
		Scan(&doctor)

	if err := qry.Error; err != nil {
		return nil, err
	}

	// Retrieve experience, education, and workday for the doctor
	experience, err := pdata.GetByIDExperience(id)
	if err != nil {
		return nil, err
	}
	education, err := pdata.GetByIDEducation(id)
	if err != nil {
		return nil, err
	}
	workday, err := pdata.GetByIDWorkadays(id)
	if err != nil {
		return nil, err
	}

	// Assign the retrieved data to the corresponding fields in the Doctor struct
	doctor.DoctorExperience = experience
	doctor.DoctorEducation = education
	doctor.DoctorWorkday = workday

	return &doctor, nil
}

func (pdata *DoctorData) GetByIDExperience(id int) ([]doctor.DoctorInfoExperience, error) {
	var doctorInfoExperience []doctor.DoctorInfoExperience

	qry := pdata.db.Table("doctors_experience").
		Select("doctors_experience.*").Where("doctors_experience.doctor_id = ?", id).Where("doctors_experience.deleted_at IS NULL").
		Scan(&doctorInfoExperience)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoExperience, nil
}
func (pdata *DoctorData) GetByIDEducation(id int) ([]doctor.DoctorInfoEducation, error) {
	var doctorInfoEducation []doctor.DoctorInfoEducation

	qry := pdata.db.Table("doctors_education").
		Select("doctors_education.*").Where("doctors_education.doctor_id = ?", id).Where("doctors_education.deleted_at IS NULL").
		Scan(&doctorInfoEducation)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoEducation, nil
}

func (pdata *DoctorData) GetByIDWorkadays(id int) ([]doctor.DoctorInfoWorkday, error) {
	var doctorInfoWorkday []doctor.DoctorInfoWorkday

	qry := pdata.db.Table("doctors_workadays").
		Select("doctors_workadays.*").Where("doctors_workadays.doctor_id = ?", id).Where("doctors_workadays.deleted_at IS NULL").
		Scan(&doctorInfoWorkday)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoWorkday, nil
}

// CREATE DATA QUERY \\

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

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

// UPDATE QUERY \\

func (pdata *DoctorData) UpdateDoctorDatapokok(id int, newData doctor.DoctorDatapokokUpdate) (bool, error) {

	var qry = pdata.db.Table("doctors").Where("id = ?", id).Updates(doctor.DoctorDatapokokUpdate{
		DoctorName:          newData.DoctorName,
		DoctorExperienced:   newData.DoctorExperienced,
		DoctorDescription:   newData.DoctorDescription,
		DoctorAvatar:        newData.DoctorAvatar,
		DoctorOfficeName:    newData.DoctorOfficeName,
		DoctorOfficeAddress: newData.DoctorOfficeAddress,
		DoctorOfficeCity:    newData.DoctorOfficeCity,
		DoctorMeetLink:      newData.DoctorMeetLink,
		DoctorSIPP:          newData.DoctorSIPP,
		DoctorSIPPFile:      newData.DoctorSIPPFile,
		DoctorSTR:           newData.DoctorSTR,
		DoctorSTRFile:       newData.DoctorSTRFile,
		DoctorCV:            newData.DoctorCV,
		DoctorIjazah:        newData.DoctorIjazah,
		DoctorExpertiseID:   newData.DoctorExpertiseID,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, nil
	}

	var qryExpertise = pdata.db.Table("doctors_expertise_relation").Where("id = ?", id).Updates(doctor.DoctorExpertiseRelation{
		ExpertiseID: newData.DoctorExpertiseID,
	})

	if err := qryExpertise.Error; err != nil {
		return false, err
	}

	if dataCount := qryExpertise.RowsAffected; dataCount < 1 {
		return false, nil
	}

	return true, nil

}

func (pdata *DoctorData) UpdateDoctorEducation(id int, doctorID int, newData doctor.DoctorInfoEducation) (bool, error) {

	var checkData doctor.DoctorInfoEducation

	qryCheck := pdata.db.Table("doctors_education").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.DoctorID) != doctorID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_education").Where("id = ?", id).Updates(doctor.DoctorInfoEducation{
		DoctorUniversity:   newData.DoctorUniversity,
		DoctorStudyProgram: newData.DoctorStudyProgram,
		DoctorGraduateYear: newData.DoctorGraduateYear,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) UpdateDoctorExperience(id int, doctorID int, newData doctor.DoctorInfoExperience) (bool, error) {
	var checkData doctor.DoctorInfoExperience

	qryCheck := pdata.db.Table("doctors_experience").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.DoctorID) != doctorID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_experience").Where("id = ?", id).Updates(doctor.DoctorInfoExperience{
		DoctorCompany:               newData.DoctorCompany,
		DoctorTitle:                 newData.DoctorTitle,
		DoctorExperienceDescription: newData.DoctorExperienceDescription,
		DoctorStartDate:             newData.DoctorStartDate,
		DoctorEndDate:               newData.DoctorEndDate,
		DoctorIsNow:                 newData.DoctorIsNow,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) UpdateDoctorWorkdays(id int, doctorID int, newData doctor.DoctorInfoWorkday) (bool, error) {

	var checkData doctor.DoctorInfoWorkday

	qryCheck := pdata.db.Table("doctors_workadays").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.DoctorID) != doctorID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_workadays").Where("id = ?", id).Updates(doctor.DoctorInfoWorkday{
		WorkdayID: newData.WorkdayID,
		StartTime: newData.StartTime,
		EndTime:   newData.EndTime,
	})

	if err := qry.Error; err != nil {
		return false, nil
	}

	return true, nil
}

func (pdata *DoctorData) getTotalConseling(id int) (int, int, int, int) {
	var totalPatient int64
	var tipe1 int64
	var tipe2 int64
	var totalJamPraktek int64
	var totalLayananChat int64
	var totalLayananVideoCall int64

	var _ = pdata.db.Table("transactions").Where("doctor_id = ? ", id).Count(&totalPatient)
	var _ = pdata.db.Table("transactions").Where("doctor_id = ? AND duration_id = ? ", id, "1").Count(&tipe1)
	var _ = pdata.db.Table("transactions").Where("doctor_id = ? AND duration_id = ? ", id, "2").Count(&tipe2)
	var _ = pdata.db.Table("transactions").Where("doctor_id = ? AND method_id = ?", id, "1").Count(&totalLayananChat)
	var _ = pdata.db.Table("transactions").Where("doctor_id = ? AND method_id = ?", id, "2").Count(&totalLayananVideoCall)

	totalJamPraktek = ((tipe1 * 60) + (tipe2 * 90))
	totalPatientInt := int(totalPatient)
	totalJamPraktekInt := int(totalJamPraktek / 60)
	totalLayananChatInt := int(totalLayananChat)
	totalLayananVideoCallInt := int(totalLayananVideoCall)

	return totalPatientInt, totalJamPraktekInt, totalLayananChatInt, totalLayananVideoCallInt
}

func (pdata *DoctorData) DoctorDashboard(id int) (doctor.DoctorDashboard, error) {
	var dashboardDoctor doctor.DoctorDashboard
	tPatient, tJamPraktek, tLayananChat, tLayananVideoCall := pdata.getTotalConseling(id)

	dashboardDoctor.TotalPatient = tPatient
	dashboardDoctor.TotalJamPraktek = tJamPraktek
	dashboardDoctor.TotalLayananChat = tLayananChat
	dashboardDoctor.TotalLayananVideoCall = tLayananVideoCall

	return dashboardDoctor, nil
}

func (pdata *DoctorData) DoctorDashboardPatient(id int) ([]doctor.DoctorDashboardPatient, error) {
	var doctors = []doctor.DoctorDashboardPatient{}

	var qry = pdata.db.Table("transactions").Select("transactions.*,patient_accounts.name as patient_name,patient_accounts.gender as gender, counseling_topics.name as topic, counseling_methods.name as Layanan").
		Joins("LEFT JOIN patient_accounts ON patient_accounts.id = transactions.patient_id").
		Joins("LEFT JOIN counseling_topics ON counseling_topics.id = transactions.topic_id").
		Joins("LEFT JOIN counseling_methods ON counseling_methods.id = transactions.method_id").
		Where("transactions.doctor_id = ?", id).
		Where("transactions.deleted_at is null").
		Scan(&doctors)

	if err := qry.Error; err != nil {
		return nil, err
	}
	return doctors, nil
}
