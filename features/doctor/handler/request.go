package handler

import "time"

type DoctorRequest struct {
	UserID              uint   `json:"user_id" form:"user_id"`
	DoctorName          string `json:"doctor_name" form:"doctor_name"`
	DoctorExperienced   string `json:"doctor_experienced" form:"doctor_experienced"`
	DoctorDescription   string `json:"doctor_description" form:"doctor_description"`
	DoctorAvatar        string `json:"doctor_avatar" form:"doctor_avatar"`
	DoctorOfficeName    string `json:"doctor_office_name" form:"doctor_office_name"`
	DoctorOfficeAddress string `json:"doctor_office_address" form:"doctor_office_address"`
	DoctorOfficeCity    string `json:"doctor_office_city" form:"doctor_office_city"`
	DoctorMeetLink      string `json:"doctor_meet_link" form:"doctor_meet_link"`
	DoctorSIPP          uint   `json:"doctor_sipp" form:"doctor_sipp"`
	DoctorSIPPFile      string `json:"doctor_sipp_file" form:"doctor_sipp_file"`
	DoctorSTR           uint   `json:"doctor_str" form:"doctor_str"`
	DoctorSTRFile       string `json:"doctor_str_file" form:"doctor_str_file"`
	DoctorCV            string `json:"doctor_cv" form:"doctor_cv"`
	DoctorIjazah        string `json:"doctor_ijazah" form:"doctor_ijazah"`
	DoctorBalance       uint   `json:"doctor_balance" form:"doctor_balance"`
	DoctorStatus        string `json:"doctor_status" form:"doctor_status"`

	//FOR DOCTOR EXPERTISE
	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id"`

	//FOR DOCTOR EDUCATION
	DoctorUniversity   []string    `json:"doctor_university" form:"doctor_university"`
	DoctorStudyProgram []string    `json:"doctor_study_program" form:"doctor_study_program"`
	DoctorGraduateYear []time.Time `json:"doctor_graduate_year" form:"doctor_graduate_year"`

	//FOR DOCTOR EXPERIENCE
	DoctorCompany               []string    `json:"doctor_company" form:"doctor_company"`
	DoctorTitle                 []string    `json:"doctor_title" form:"doctor_title"`
	DoctorExperienceDescription []string    `json:"doctor_experience_description" form:"doctor_experience_description"`
	DoctorStartDate             []time.Time `json:"doctor_start_date" form:"doctor_start_date"`
	DoctorEndDate               []time.Time `json:"doctor_end_date" form:"doctor_end_date"`
	DoctorIsNow                 []bool      `json:"doctor_is_now" form:"doctor_is_now"`

	//FOR DOCTOR WORKDAY
	DoctorWorkdayID     []uint      `json:"workday_id" form:"workday_id"`
	DoctorWorkStartTime []time.Time `json:"start_time" form:"start_time"`
	DoctorWorkEndTime   []time.Time `json:"end_time" form:"end_time"`
}
