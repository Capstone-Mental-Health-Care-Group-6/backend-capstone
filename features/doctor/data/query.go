package data

import (
	"FinalProject/features/doctor"
	"FinalProject/features/users"
	"errors"
	"fmt"
	"time"

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
		Joins("LEFT JOIN doctors_expertise_relation ON doctors_expertise_relation.doctor_id = doctors.id").
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
		rating, err := pdata.GetByIDRating(int(doctor.ID))
		if err != nil {
			return nil, err
		}

		// Assign the retrieved data to the corresponding fields in the Doctor struct
		doctors[i].DoctorExperience = experience
		doctors[i].DoctorEducation = education
		doctors[i].DoctorWorkdays = workday
		doctors[i].DoctorRating = rating
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

func (pdata *DoctorData) GetAll(name string) ([]doctor.DoctorAll, error) {
	var doctors []doctor.DoctorAll

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors_expertise_relation.doctor_id = doctors.id").
		Where("doctors.deleted_at IS NULL")

	if name != "" {
		qry.Where("doctors.doctor_name LIKE ?", "%"+name+"%")
	}

	if err := qry.Scan(&doctors).Error; err != nil {
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
		rating, err := pdata.GetByIDRating(int(doctor.ID))
		if err != nil {
			return nil, err
		}

		fmt.Println("Data", doctor.ID)

		// Assign the retrieved data to the corresponding fields in the Doctor struct
		doctors[i].DoctorExperience = experience
		doctors[i].DoctorEducation = education
		doctors[i].DoctorWorkdays = workday
		doctors[i].DoctorRating = rating
	}

	return doctors, nil
}

func (pdata *DoctorData) GetByID(id int) (*doctor.DoctorAll, error) {
	var doctor doctor.DoctorAll

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors_expertise_relation.doctor_id = doctors.id").
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
	rating, err := pdata.GetByIDRating(int(doctor.ID))
	if err != nil {
		return nil, err
	}

	// Assign the retrieved data to the corresponding fields in the Doctor struct
	doctor.DoctorExperience = experience
	doctor.DoctorEducation = education
	doctor.DoctorWorkdays = workday
	doctor.DoctorRating = rating

	return &doctor, nil
}

func (pdata *DoctorData) GetDoctorByUserId(userId int) (*doctor.DoctorAll, error) {
	var doctor doctor.DoctorAll

	qry := pdata.db.Table("doctors").
		Select("doctors.*, doctors_expertise_relation.expertise_id AS expertise_id").
		Joins("LEFT JOIN doctors_expertise_relation ON doctors_expertise_relation.doctor_id = doctors.id").
		Where("doctors.user_id = ?", userId).
		Where("doctors.deleted_at IS NULL").
		Scan(&doctor)

	if err := qry.Error; err != nil {
		return nil, err
	}

	// Retrieve experience, education, and workday for the doctor
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

	rating, err := pdata.GetByIDRating(int(doctor.ID))
	if err != nil {
		return nil, err
	}

	// Assign the retrieved data to the corresponding fields in the Doctor struct
	doctor.DoctorExperience = experience
	doctor.DoctorEducation = education
	doctor.DoctorWorkdays = workday
	doctor.DoctorRating = rating

	return &doctor, nil
}

func (pdata *DoctorData) GetByIDExperience(id int) ([]doctor.DoctorExperience, error) {
	var doctorInfoExperience []doctor.DoctorExperience

	qry := pdata.db.Table("doctors_experience").
		Select("doctors_experience.*").Where("doctors_experience.doctor_id = ?", id).Where("doctors_experience.deleted_at IS NULL").
		Scan(&doctorInfoExperience)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoExperience, nil
}
func (pdata *DoctorData) GetByIDEducation(id int) ([]doctor.DoctorEducation, error) {
	var doctorInfoEducation []doctor.DoctorEducation

	qry := pdata.db.Table("doctors_education").
		Select("doctors_education.*").Where("doctors_education.doctor_id = ?", id).Where("doctors_education.deleted_at IS NULL").
		Scan(&doctorInfoEducation)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoEducation, nil
}

func (pdata *DoctorData) GetByIDWorkadays(id int) ([]doctor.DoctorWorkdays, error) {
	var doctorInfoWorkday []doctor.DoctorWorkdays

	qry := pdata.db.Table("doctors_workadays").
		Select("doctors_workadays.*").Where("doctors_workadays.doctor_id = ?", id).Where("doctors_workadays.deleted_at IS NULL").
		Scan(&doctorInfoWorkday)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoWorkday, nil
}

// func (pdata *DoctorData) GetByIDRating(id int) ([]doctor.DoctorRating, error) {
// 	var doctorInfoRating []doctor.DoctorRating

// 	qry := pdata.db.Table("doctors_rating").
// 		Select("doctors_rating.*").Where("doctors_rating.doctor_id = ?", id).Where("doctors_rating.deleted_at IS NULL").
// 		Scan(&doctorInfoRating)

// 	if err := qry.Error; err != nil {
// 		return nil, err
// 	}

// 	return doctorInfoRating, nil
// }

func (pdata *DoctorData) GetByIDRating(id int) ([]doctor.DoctorRating, error) {
	var doctorInfoDetailRating []doctor.DoctorRating

	qry := pdata.db.Table("doctors_rating").
		Select(`doctors_rating.*, 
		patient_accounts.name as patient_name,
        patient_accounts.avatar as patient_avatar,
        doctors.doctor_name as doctor_name,
        doctors.doctor_avatar as doctor_avatar`).
		Joins("LEFT JOIN patient_accounts ON patient_accounts.id = doctors_rating.patient_id").
		Joins("LEFT JOIN doctors ON doctors.id = doctors_rating.doctor_id").
		Where("doctors_rating.doctor_id = ?", id).
		Where("doctors_rating.deleted_at IS NULL").
		Scan(&doctorInfoDetailRating)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return doctorInfoDetailRating, nil
}

// CREATE DATA QUERY \\

func (pdata *DoctorData) Insert(newData doctor.Doctor) (*doctor.Doctor, error) {

	var dbData = new(Doctor)
	dbData.DoctorName = newData.DoctorName
	dbData.UserID = newData.UserID
	dbData.DoctorNIK = newData.DoctorNIK
	dbData.DoctorDOB = newData.DoctorDOB
	dbData.DoctorProvinsi = newData.DoctorProvinsi
	dbData.DoctorKota = newData.DoctorKota
	dbData.DoctorNumberPhone = newData.DoctorNumberPhone
	dbData.DoctorGender = newData.DoctorGender
	dbData.DoctorAvatar = newData.DoctorAvatar
	dbData.DoctorDescription = newData.DoctorDescription
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

func (pdata *DoctorData) InsertWorkadays(newData doctor.DoctorWorkdays) (*doctor.DoctorWorkdays, error) {

	var dbData = new(DoctorWorkadays)
	dbData.DoctorID = newData.DoctorID
	dbData.WorkdayID = newData.WorkdayID
	dbData.StartTime = newData.StartTime
	dbData.EndTime = newData.EndTime

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	newData.ID = dbData.ID

	return &newData, nil
}

func (pdata *DoctorData) InsertExperience(newData doctor.DoctorExperience) (*doctor.DoctorExperience, error) {

	var dbData = new(DoctorExperience)
	dbData.DoctorID = newData.DoctorID
	dbData.DoctorCompany = newData.DoctorCompany
	dbData.DoctorTitle = newData.DoctorTitle
	dbData.DoctorCompanyAddress = newData.DoctorCompany
	dbData.DoctorStartDate = newData.DoctorStartDate
	dbData.DoctorEndDate = newData.DoctorEndDate

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	newData.ID = dbData.ID

	return &newData, nil
}

func (pdata *DoctorData) InsertEducation(newData doctor.DoctorEducation) (*doctor.DoctorEducation, error) {

	var dbData = new(DoctorEducation)
	dbData.DoctorID = newData.DoctorID
	dbData.DoctorUniversity = newData.DoctorUniversity
	dbData.DoctorStudyProgram = newData.DoctorStudyProgram
	dbData.DoctorEnrollYear = newData.DoctorEnrollYear
	dbData.DoctorGraduateYear = newData.DoctorGraduateYear

	if err := pdata.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	newData.ID = dbData.ID

	return &newData, nil
}

func (pdata *DoctorData) IsLinkUsed(meetLink string) bool {
	var count int64
	pdata.db.Table("doctors").Where("doctor_meet_link = ?", meetLink).Count(&count)
	return count > 0
}

// UPDATE QUERY \\

func (pdata *DoctorData) UpdateDoctorDatapokok(id int, newData doctor.DoctorDatapokokUpdate) (bool, error) {

	var qry = pdata.db.Table("doctors").Where("id = ?", id).Updates(doctor.DoctorDatapokokUpdate{
		DoctorName:        newData.DoctorName,
		UserID:            newData.UserID,
		DoctorNIK:         newData.DoctorNIK,
		DoctorProvinsi:    newData.DoctorProvinsi,
		DoctorKota:        newData.DoctorKota,
		DoctorNumberPhone: newData.DoctorNumberPhone,
		DoctorGender:      newData.DoctorGender,
		DoctorAvatar:      newData.DoctorAvatar,
		DoctorDescription: newData.DoctorDescription,
		DoctorMeetLink:    newData.DoctorMeetLink,
		DoctorSIPP:        newData.DoctorSIPP,
		DoctorSIPPFile:    newData.DoctorSIPPFile,
		DoctorSTR:         newData.DoctorSTR,
		DoctorSTRFile:     newData.DoctorSTRFile,
		DoctorCV:          newData.DoctorCV,
		DoctorIjazah:      newData.DoctorIjazah,
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

func (pdata *DoctorData) UpdateDoctorEducation(id int, doctorID int, newData doctor.DoctorEducation) (bool, error) {

	var checkData doctor.DoctorEducation

	qryCheck := pdata.db.Table("doctors_education").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.DoctorID) != doctorID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_education").Where("id = ?", id).Updates(doctor.DoctorEducation{
		DoctorUniversity:   newData.DoctorUniversity,
		DoctorStudyProgram: newData.DoctorStudyProgram,
		DoctorGraduateYear: newData.DoctorGraduateYear,
		DoctorEnrollYear:   newData.DoctorEnrollYear,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) UpdateDoctorExperience(id int, doctorID int, newData doctor.DoctorExperience) (bool, error) {
	var checkData doctor.DoctorExperience

	qryCheck := pdata.db.Table("doctors_experience").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.DoctorID) != doctorID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_experience").Where("id = ?", id).Updates(doctor.DoctorExperience{
		DoctorCompany:        newData.DoctorCompany,
		DoctorTitle:          newData.DoctorTitle,
		DoctorCompanyAddress: newData.DoctorCompanyAddress,
		DoctorStartDate:      newData.DoctorStartDate,
		DoctorEndDate:        newData.DoctorEndDate,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) UpdateDoctorWorkdays(id int, doctorID int, newData doctor.DoctorWorkdays) (bool, error) {

	var checkData doctor.DoctorWorkdays

	qryCheck := pdata.db.Table("doctors_workadays").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.DoctorID) != doctorID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_workadays").Where("id = ?", id).Updates(doctor.DoctorWorkdays{
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

func (pdata *DoctorData) UpdateDoctorRating(id int, patientID int, newData doctor.DoctorRating) (bool, error) {

	var checkData doctor.DoctorRating

	qryCheck := pdata.db.Table("doctors_rating").Where("id = ?", id).First(&checkData)

	if err := qryCheck.Error; err != nil {
		return false, nil
	}

	if int(checkData.PatientID) != patientID {
		return false, nil
	}

	qry := pdata.db.Table("doctors_rating").Where("id = ?", id).Updates(doctor.DoctorRating{
		DoctorStarRating: newData.DoctorStarRating,
		DoctorReview:     newData.DoctorReview,
	})

	if err := qry.Error; err != nil {
		return false, nil
	}

	return true, nil
}

func (pdata *DoctorData) DeleteDoctorWorkdays(doctorID int) (bool, error) {
	var deleteData = new(doctor.DoctorWorkdays)

	if err := pdata.db.Table("doctors_workadays").Where("id = ?", doctorID).Delete(deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) DeleteDoctorExperience(doctorID int) (bool, error) {
	var deleteData = new(doctor.DoctorExperience)

	if err := pdata.db.Table("doctors_experience").Where("id = ?", doctorID).Delete(deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) DeleteDoctorEducation(doctorID int) (bool, error) {
	var deleteData = new(doctor.DoctorEducation)

	if err := pdata.db.Table("doctors_education").Where("id = ?", doctorID).Delete(deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) DeleteDoctorRating(doctorID int) (bool, error) {
	var deleteData = new(doctor.DoctorRating)

	if err := pdata.db.Table("doctors_education").Where("id = ?", doctorID).Delete(deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) DeleteDoctor(doctorID int) (bool, error) {
	var deleteData = new(doctor.DoctorAll)

	if err := pdata.db.Table("doctors_rating").Where("id = ?", doctorID).Delete(deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pdata *DoctorData) DoctorDashboardAdmin() (doctor.DoctorDashboardAdmin, error) {
	var dashboardArticle doctor.DoctorDashboardAdmin

	tDokter, tDokterBaru, tDokterPending, tDokterActive := pdata.getTotalDoctor()

	dashboardArticle.TotalDoctor = tDokter
	dashboardArticle.TotalDoctorActive = tDokterActive
	dashboardArticle.TotalDoctorBaru = tDokterBaru
	dashboardArticle.TotalDoctorPending = tDokterPending

	return dashboardArticle, nil
}

func (pdata *DoctorData) getTotalDoctor() (int, int, int, int) {
	var totalDokter int64
	var totalDokterBaru int64
	var totalDokterPending int64
	var totalDokterAktif int64

	var now = time.Now()
	var before = now.AddDate(0, 0, -30)

	var _ = pdata.db.Table("doctors").Count(&totalDokter)
	var _ = pdata.db.Table("doctors").Where("created_at BETWEEN ? and ?", before, now).Where("doctor_status = ?", "Confirmed").Count(&totalDokterBaru)
	var _ = pdata.db.Table("doctors").Where("doctor_status = ?", "Request").Count(&totalDokterPending)
	var _ = pdata.db.Table("doctors").Where("doctor_status = ?", "Confirmed").Count(&totalDokterAktif)

	totalDokterInt := int(totalDokter)
	totalDokterBaruInt := int(totalDokterBaru)
	totalDokterPendingInt := int(totalDokterPending)
	totalDokterAktifInt := int(totalDokterAktif)

	return totalDokterInt, totalDokterBaruInt, totalDokterPendingInt, totalDokterAktifInt
}

func (pdata *DoctorData) DenyDoctor(id int) (bool, error) {
	var qry = pdata.db.Table("doctors").Where("id = ?", id).Updates(Doctor{DoctorStatus: "Reject"})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (pdata *DoctorData) ApproveDoctor(id int) (bool, error) {
	var qry = pdata.db.Table("doctors").Where("id = ?", id).Updates(Doctor{DoctorStatus: "Confirmed"})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
