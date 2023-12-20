package data

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	*gorm.Model
	UserID     uint   `gorm:"column:user_id"`
	DoctorName string `gorm:"column:doctor_name;type:varchar(255)"`

	DoctorNIK         string `gorm:"column:doctor_nik;type:varchar(255);unique"`
	DoctorDOB         string `gorm:"column:doctor_dob;type:varchar(255)"`
	DoctorProvinsi    string `gorm:"column:doctor_provinsi;type:varchar(255)"`
	DoctorKota        string `gorm:"column:doctor_kota;type:varchar(255)"`
	DoctorNumberPhone string `gorm:"column:doctor_number_phone;type:varchar(255)"`
	DoctorGender      string `gorm:"column:doctor_gender;type:enum('laki','perempuan')"`
	DoctorAvatar      string `gorm:"column:doctor_avatar;type:varchar(255)"`
	DoctorDescription string `gorm:"column:doctor_description;type:varchar(255)"`

	DoctorMeetLink          string                    `gorm:"column:doctor_meet_link;type:varchar(255)"`
	DoctorSIPP              string                    `gorm:"column:doctor_sipp;type:varchar(255)"`
	DoctorSIPPFile          string                    `gorm:"column:doctor_sipp_file;type:varchar(255)"`
	DoctorSTR               string                    `gorm:"column:doctor_str;type:varchar(255)"`
	DoctorSTRFile           string                    `gorm:"column:doctor_str_file;type:varchar(255)"`
	DoctorCV                string                    `gorm:"column:doctor_cv;type:varchar(255)"`
	DoctorIjazah            string                    `gorm:"column:doctor_ijazah;type:varchar(255)"`
	DoctorBalance           uint                      `gorm:"column:doctor_balance"`
	DoctorStatus            string                    `gorm:"column:doctor_status;type:enum('Request','Confirmed','Reject')"`
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
	DoctorEnrollYear   time.Time `gorm:"column:doctor_enroll_year"`
	DoctorGraduateYear time.Time `gorm:"column:doctor_graduate_year"`
}

type DoctorExperience struct {
	*gorm.Model
	DoctorID             uint      `gorm:"column:doctor_id;index"`
	DoctorCompany        string    `gorm:"column:doctor_company"`
	DoctorTitle          string    `gorm:"column:doctor_title"`
	DoctorCompanyAddress string    `gorm:"column:doctor_company_address"`
	DoctorStartDate      time.Time `gorm:"column:doctor_start_date"`
	DoctorEndDate        time.Time `gorm:"column:doctor_end_date"`
	// DoctorIsNow     bool      `gorm:"column:doctor_is_now"`
}

type DoctorRating struct {
	*gorm.Model
	DoctorID         uint   `gorm:"column:doctor_id;index"`
	PatientID        uint   `gorm:"column:patient_id"`
	TransactionID    string `gorm:"column:transaction_id"`
	DoctorStarRating uint   `gorm:"column:doctor_star_rating"`
	DoctorReview     string `gorm:"column:doctor_review;type:varchar(255)"`
}

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
