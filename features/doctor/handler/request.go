package handler

import "time"

type DoctorRequest struct {
	UserID            uint   `json:"user_id" form:"user_id"`
	DoctorName        string `json:"doctor_name" form:"doctor_name"`
	DoctorNIK         string `json:"doctor_nik" form:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob" form:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi" form:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota" form:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_number_phone" form:"doctor_number_phone"`
	DoctorGender      string `json:"doctor_gender" form:"doctor_gender"`
	DoctorAvatar      string `json:"doctor_avatar" form:"doctor_avatar"`
	DoctorDescription string `json:"doctor_description" form:"doctor_description"`
	DoctorMeetLink    string `json:"doctor_meet_link" form:"doctor_meet_link"`
	DoctorSIPP        string `json:"doctor_sipp" form:"doctor_sipp"`
	DoctorSIPPFile    string `json:"doctor_sipp_file" form:"doctor_sipp_file"`
	DoctorSTR         string `json:"doctor_str" form:"doctor_str"`
	DoctorSTRFile     string `json:"doctor_str_file" form:"doctor_str_file"`
	DoctorCV          string `json:"doctor_cv" form:"doctor_cv"`
	DoctorIjazah      string `json:"doctor_ijazah" form:"doctor_ijazah"`
	DoctorBalance     uint   `json:"doctor_balance" form:"doctor_balance"`
	DoctorStatus      string `json:"doctor_status" form:"doctor_status"`

	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id"`

	//FOR DOCTOR EDUCATION
	DoctorUniversity   []string    `json:"doctor_university" form:"doctor_university"`
	DoctorStudyProgram []string    `json:"doctor_study_program" form:"doctor_study_program"`
	DoctorEnrollYear   []time.Time `json:"doctor_enroll_year" form:"doctor_enroll_year"`
	DoctorGraduateYear []time.Time `json:"doctor_graduate_year" form:"doctor_graduate_year"`

	//FOR DOCTOR EXPERIENCE
	DoctorCompany        []string    `json:"doctor_company" form:"doctor_company"`
	DoctorTitle          []string    `json:"doctor_title" form:"doctor_title"`
	DoctorStartDate      []time.Time `json:"doctor_start_date" form:"doctor_start_date"`
	DoctorEndDate        []time.Time `json:"doctor_end_date" form:"doctor_end_date"`
	DoctorCompanyAddress []string    `json:"doctor_company_address" form:"doctor_company_address"`

	//FOR DOCTOR WORKDAY
	DoctorWorkdayID     []uint      `json:"workday_id" form:"workday_id"`
	DoctorWorkStartTime []time.Time `json:"start_time" form:"start_time"`
	DoctorWorkEndTime   []time.Time `json:"end_time" form:"end_time"`
}

type DoctorRequestDatapokok struct {
	UserID            uint   `json:"user_id" form:"user_id"`
	DoctorName        string `json:"doctor_name" form:"doctor_name"`
	DoctorNIK         string `json:"doctor_nik" form:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob" form:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi" form:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota" form:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_number_phone" form:"doctor_number_phone"`
	DoctorGender      string `json:"doctor_gender" form:"doctor_gender"`
	DoctorAvatar      string `json:"doctor_avatar" form:"doctor_avatar"`
	DoctorDescription string `json:"doctor_description" form:"doctor_description"`
	DoctorMeetLink    string `json:"doctor_meet_link" form:"doctor_meet_link"`
	DoctorSIPP        string `json:"doctor_sipp" form:"doctor_sipp"`
	DoctorSIPPFile    string `json:"doctor_sipp_file" form:"doctor_sipp_file"`
	DoctorSTR         string `json:"doctor_str" form:"doctor_str"`
	DoctorSTRFile     string `json:"doctor_str_file" form:"doctor_str_file"`
	DoctorCV          string `json:"doctor_cv" form:"doctor_cv"`
	DoctorIjazah      string `json:"doctor_ijazah" form:"doctor_ijazah"`
	DoctorBalance     uint   `json:"doctor_balance" form:"doctor_balance"`
	DoctorStatus      string `json:"doctor_status" form:"doctor_status"`

	//FOR DOCTOR EXPERTISE
	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id"`
}

type DoctorWorkdays struct {
	ID        uint      `json:"id"`
	DoctorID  uint      `json:"doctor_id"`
	WorkdayID uint      `json:"workday_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type DoctorEducation struct {
	ID                 uint      `json:"id"`
	DoctorID           uint      `json:"doctor_id"`
	DoctorUniversity   string    `json:"doctor_university"`
	DoctorStudyProgram string    `json:"doctor_study_program"`
	DoctorEnrollYear   time.Time `json:"doctor_enroll_year"`
	DoctorGraduateYear time.Time `json:"doctor_graduate_year"`
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
