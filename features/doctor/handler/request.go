package handler

import "time"

type DoctorRequest struct {
	UserID uint `json:"user_id" form:"user_id" validate:"required"`

	DoctorName string `json:"doctor_name" form:"doctor_name" validate:"required"`
	DoctorNIK  string `json:"doctor_nik" form:"doctor_nik" validate:"required"`
	DoctorDOB  string `json:"doctor_dob" form:"doctor_dob" validate:"required"`

	DoctorProvinsi string `json:"doctor_provinsi" form:"doctor_provinsi" validate:"required"`
	DoctorKota     string `json:"doctor_kota" form:"doctor_kota" validate:"required"`

	DoctorNumberPhone string `json:"doctor_number_phone" form:"doctor_number_phone" validate:"required"`
	DoctorGender      string `json:"doctor_gender" form:"doctor_gender" validate:"required"`
	DoctorAvatar      string `json:"doctor_avatar" form:"doctor_avatar" validate:"required"`
	DoctorDescription string `json:"doctor_description" form:"doctor_description" validate:"required"`

	DoctorSIPP string `json:"doctor_sipp" form:"doctor_sipp" validate:"required"`
	DoctorSTR  string `json:"doctor_str" form:"doctor_str" validate:"required"`

	DoctorSIPPFile string `json:"doctor_sipp_file" form:"doctor_sipp_file" validate:"required"`
	DoctorSTRFile  string `json:"doctor_str_file" form:"doctor_str_file" validate:"required"`
	DoctorCV       string `json:"doctor_cv" form:"doctor_cv" validate:"required"`
	DoctorIjazah   string `json:"doctor_ijazah" form:"doctor_ijazah" validate:"required"`

	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id" validate:"required"`

	//FOR DOCTOR EDUCATION
	DoctorUniversity   []string    `json:"doctor_university" form:"doctor_university" validate:"required"`
	DoctorStudyProgram []string    `json:"doctor_study_program" form:"doctor_study_program" validate:"required"`
	DoctorEnrollYear   []time.Time `json:"doctor_enroll_year" form:"doctor_enroll_year" validate:"required"`
	DoctorGraduateYear []time.Time `json:"doctor_graduate_year" form:"doctor_graduate_year" validate:"required"`

	//FOR DOCTOR EXPERIENCE
	DoctorCompany        []string    `json:"doctor_company" form:"doctor_company" validate:"required"`
	DoctorTitle          []string    `json:"doctor_title" form:"doctor_title" validate:"required"`
	DoctorStartDate      []time.Time `json:"doctor_start_date" form:"doctor_start_date" validate:"required"`
	DoctorEndDate        []time.Time `json:"doctor_end_date" form:"doctor_end_date" validate:"required"`
	DoctorCompanyAddress []string    `json:"doctor_company_address" form:"doctor_company_address" validate:"required"`

	//FOR DOCTOR WORKDAY
	DoctorWorkdayID     []uint      `json:"workday_id" form:"workday_id" validate:"required"`
	DoctorWorkStartTime []time.Time `json:"start_time" form:"start_time" validate:"required"`
	DoctorWorkEndTime   []time.Time `json:"end_time" form:"end_time" validate:"required"`
}

type DoctorRequestDatapokok struct {
	UserID            uint   `json:"user_id" form:"user_id" validate:"required"`
	DoctorName        string `json:"doctor_name" form:"doctor_name" validate:"required"`
	DoctorNIK         string `json:"doctor_nik" form:"doctor_nik" validate:"required"`
	DoctorDOB         string `json:"doctor_dob" form:"doctor_dob" validate:"required"`
	DoctorProvinsi    string `json:"doctor_provinsi" form:"doctor_provinsi" validate:"required"`
	DoctorKota        string `json:"doctor_kota" form:"doctor_kota" validate:"required"`
	DoctorNumberPhone string `json:"doctor_number_phone" form:"doctor_number_phone" validate:"required"`
	DoctorGender      string `json:"doctor_gender" form:"doctor_gender" validate:"required"`
	DoctorAvatar      string `json:"doctor_avatar" form:"doctor_avatar" validate:"required"`
	DoctorDescription string `json:"doctor_description" form:"doctor_description" validate:"required"`
	DoctorMeetLink    string `json:"doctor_meet_link" form:"doctor_meet_link" validate:"required"`
	DoctorSIPP        string `json:"doctor_sipp" form:"doctor_sipp" validate:"required"`
	DoctorSIPPFile    string `json:"doctor_sipp_file" form:"doctor_sipp_file" validate:"required"`
	DoctorSTR         string `json:"doctor_str" form:"doctor_str" validate:"required"`
	DoctorSTRFile     string `json:"doctor_str_file" form:"doctor_str_file" validate:"required"`
	DoctorCV          string `json:"doctor_cv" form:"doctor_cv" validate:"required"`
	DoctorIjazah      string `json:"doctor_ijazah" form:"doctor_ijazah" validate:"required"`
	DoctorBalance     uint   `json:"doctor_balance" form:"doctor_balance" validate:"required"`
	DoctorStatus      string `json:"doctor_status" form:"doctor_status" validate:"required"`

	//FOR DOCTOR EXPERTISE
	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id" validate:"required"`
}

type DoctorWorkdays struct {
	ID        uint      `json:"id"`
	DoctorID  uint      `json:"doctor_id"`
	WorkdayID uint      `json:"workday_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
}

type DoctorEducation struct {
	ID                 uint      `json:"id" `
	DoctorID           uint      `json:"doctor_id"`
	DoctorUniversity   string    `json:"doctor_university" validate:"required"`
	DoctorStudyProgram string    `json:"doctor_study_program" validate:"required"`
	DoctorEnrollYear   time.Time `json:"doctor_enroll_year" validate:"required"`
	DoctorGraduateYear time.Time `json:"doctor_graduate_year" validate:"required"`
}

type DoctorExperience struct {
	ID                   uint      `json:"id"`
	DoctorID             uint      `json:"doctor_id"`
	DoctorCompany        string    `json:"doctor_company"`
	DoctorTitle          string    `json:"doctor_title"`
	DoctorCompanyAddress string    `json:"doctor_company_address"`
	DoctorStartDate      time.Time `json:"doctor_start_date"`
	DoctorEndDate        time.Time `json:"doctor_end_date"`
}

type DoctorRating struct {
	DoctorStarRating uint   `json:"doctor_star_rating"`
	DoctorReview     string `json:"doctor_review"`
}

type ManagePatient struct {
	DoctorID     uint   `json:"doctor_id"`
	PatientID    uint   `json:"patient_id"`
	Alasan       string `json:"alasan"`
	DetailAlasan string `json:"detail_alasan"`
}
