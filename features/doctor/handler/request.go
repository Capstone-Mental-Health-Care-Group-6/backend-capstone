package handler

import "time"

type DoctorRequest struct {
	UserID              uint   `json:"user_id" form:"user_id"`
	DoctorName          string `json:"doctor_name" form:"doctor_name"`
	DoctorExperience    string `json:"doctor_experience" form:"doctor_experience"`
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

	//INI BUAT DIA CREATE TABLE LAIN
	DoctorWorkdayID     uint      `json:"workday_id" form:"workday_id"`
	DoctorExpertiseID   uint      `json:"expertise_id" form:"expertise_id"`
	DoctorWorkStartTime time.Time `json:"start_time" form:"start_time"`
	DoctorWorkEndTime   time.Time `json:"end_time" form:"end_time"`
}
