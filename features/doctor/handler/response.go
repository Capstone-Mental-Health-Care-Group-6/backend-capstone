package handler

import "FinalProject/features/doctor"

type DoctorResponse struct {
	UserID            uint   `json:"user_id"`
	DoctorName        string `json:"doctor_name"`
	DoctorDescription string `json:"doctor_description"`
	DoctorAvatar      string `json:"doctor_avatar"`

	DoctorNIK         string `json:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_numberphone"`
	DoctorGender      string `json:"doctor_gender"`

	DoctorMeetLink  string `json:"doctor_meet_link"`
	DoctorSIPPFile  string `json:"doctor_sipp_file"`
	DoctorSTRFile   string `json:"doctor_str_file"`
	DoctorCV        string `json:"doctor_cv"`
	DoctorIjazah    string `json:"doctor_ijazah"`
	DoctorBalance   uint   `json:"doctor_balance"`
	DoctorStatus    string `json:"doctor_status"`
	DoctorExpertise uint   `json:"doctor_expertise"`

	DoctorWorkday    []*doctor.DoctorWorkdays   `json:"workday"`
	DoctorEducation  []*doctor.DoctorEducation  `json:"education"`
	DoctorExperience []*doctor.DoctorExperience `json:"experience"`
}

type UpdateResponse struct {
	ID          uint   `json:"id"`
	Status      bool   `json:"status_updated"`
	Description string `json:"description"`
}
