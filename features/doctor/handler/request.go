package handler

import "mime/multipart"

type DoctorRequest struct {
	UserID uint `json:"user_id" form:"user_id"`

	DoctorName          string         `json:"doctor_name" form:"doctor_name"`
	DoctorExperience    string         `json:"doctor_experience" form:"doctor_experience"`
	DoctorDescription   string         `json:"doctor_description" form:"doctor_description"`
	DoctorAvatar        multipart.File `json:"doctor_avatar" form:"doctor_avatar"`
	DoctorOfficeName    string         `json:"doctor_office_name" form:"doctor_office_name"`
	DoctorOfficeAddress string         `json:"doctor_office_address" form:"doctor_office_address"`
	DoctorOfficeCity    string         `json:"doctor_office_city" form:"doctor_office_city"`

	DoctorMeetLink string         `json:"doctor_meet_link" form:"doctor_meet_link"`
	DoctorSIPPFile multipart.File `json:"doctor_sipp_file" form:"doctor_sipp_file"`
	DoctorSTRFile  multipart.File `json:"doctor_str_file" form:"doctor_str_file"`
	DoctorCV       multipart.File `json:"doctor_cv" form:"doctor_cv"`
	DoctorIjazah   string         `json:"doctor_ijazah" form:"doctor_ijazah"`
}
