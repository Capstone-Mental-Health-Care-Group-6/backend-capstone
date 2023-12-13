package data

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	*gorm.Model
	UserID                  uint                      `gorm:"column:user_id"`
	DoctorName              string                    `gorm:"column:doctor_name;type:varchar(255)"`
	DoctorExperienced       string                    `gorm:"column:doctor_experienced;type:enum('under_five_years','five_to_ten_years','above_ten_years')"`
	DoctorDescription       string                    `gorm:"column:doctor_description;type:varchar(255)"`
	DoctorAvatar            string                    `gorm:"column:doctor_avatar;type:varchar(255)"`
	DoctorOfficeName        string                    `gorm:"column:doctor_office_name;type:varchar(255)"`
	DoctorOfficeAddress     string                    `gorm:"column:doctor_office_address;type:varchar(255)"`
	DoctorOfficeCity        string                    `gorm:"column:doctor_office_city;type:varchar(255)"`
	DoctorMeetLink          string                    `gorm:"column:doctor_meet_link;type:varchar(255)"`
	DoctorSIPP              uint                      `gorm:"column:doctor_sipp"`
	DoctorSIPPFile          string                    `gorm:"column:doctor_sipp_file;type:varchar(255)"`
	DoctorSTR               uint                      `gorm:"column:doctor_str"`
	DoctorSTRFile           string                    `gorm:"column:doctor_str_file;type:varchar(255)"`
	DoctorCV                string                    `gorm:"column:doctor_cv;type:varchar(255)"`
	DoctorIjazah            string                    `gorm:"column:doctor_ijazah;type:varchar(255)"`
	DoctorBalance           uint                      `gorm:"column:doctor_balance"`
	DoctorStatus            string                    `gorm:"column:doctor_status;type:enum('request','confirmed')"`
	DoctorExpertiseRelation []DoctorExpertiseRelation `gorm:"foreignKey:DoctorID"`
	DoctorWorkadays         []DoctorWorkadays         `gorm:"foreignKey:DoctorID"`
	DoctorExperience        []DoctorExperience        `gorm:"foreignKey:DoctorID"`
	DoctorEducation         []DoctorEducation         `gorm:"foreignKey:DoctorID"`
	DoctorRating            []DoctorRating            `gorm:"foreignKey:DoctorID"`
}

type DoctorExpertiseRelation struct {
	*gorm.Model
	DoctorID    uint `gorm:"column:doctor_id;index"`
	ExpertiseID uint `gorm:"column:expertise_id"`
}

type DoctorWorkadays struct {
	*gorm.Model
	DoctorID  uint      `gorm:"column:doctor_id;index"`
	WorkdayID uint      `gorm:"column:workday_id"`
	StartTime time.Time `gorm:"default:null;column:start_time"`
	EndTime   time.Time `gorm:"default:null;column:end_time"`
}

type DoctorEducation struct {
	*gorm.Model
	DoctorID           uint      `gorm:"column:doctor_id;index"`
	DoctorUniversity   string    `gorm:"column:doctor_university"`
	DoctorStudyProgram string    `gorm:"column:doctor_study_program"`
	DoctorGraduateYear time.Time `gorm:"column:doctor_graduate_year"`
}

type DoctorExperience struct {
	*gorm.Model
	DoctorID                    uint      `gorm:"column:doctor_id;index"`
	DoctorCompany               string    `gorm:"column:doctor_company"`
	DoctorTitle                 string    `gorm:"column:doctor_title"`
	DoctorExperienceDescription string    `gorm:"column:doctor_experience_description"`
	DoctorStartDate             time.Time `gorm:"column:doctor_start_date"`
	DoctorEndDate               time.Time `gorm:"column:doctor_end_date"`
	DoctorIsNow                 bool      `gorm:"column:doctor_is_now"`
}

type DoctorRating struct {
	*gorm.Model
	DoctorID         uint   `gorm:"column:doctor_id;index"`
	PatientID        uint   `gorm:"column:patient_id"`
	DoctorStarRating uint   `gorm:"column:doctor_star_rating"`
	DoctorReview     string `gorm:"column:doctor_review;type:varchar(255)"`
}

//type ManagePatient struct {
//	*gorm.Model
//	DoctorID     uint   `gorm:"column:doctor_id"`
//	PatientID    uint   `gorm:"column:patient_id"`
//	Alasan       string `gorm:"column:alasan;type:enum('Overbook','Time_limit','The_doctor_has_other_activities')"`
//	DetailAlasan string `gorm:"column:detail_alasan"`
//}

func (Doctor) TableName() string {
	return "doctors"
}

func (DoctorExpertiseRelation) TableName() string {
	return "doctors_expertise_relation"
}

func (DoctorWorkadays) TableName() string {
	return "doctors_workadays"
}

func (DoctorRating) TableName() string {
	return "doctors_rating"
}

func (DoctorEducation) TableName() string {
	return "doctors_education"
}

func (DoctorExperience) TableName() string {
	return "doctors_experience"
}

//func (ManagePatient) TableName() string {
//	return "manage_patient"
//}
