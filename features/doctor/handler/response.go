package handler

import (
	"FinalProject/features/doctor"
)

type DoctorResponse struct {
	UserID            uint   `json:"user_id"`
	DoctorName        string `json:"doctor_name"`
	DoctorDescription string `json:"doctor_description"`
	DoctorAvatar      string `json:"doctor_avatar"`

	DoctorNIK         string `json:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_number_phone"`
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

type DashboardResponse struct {
	TotalPatient          int `json:"total_patient"`
	TotalJamPraktek       int `json:"total_jam_praktek"`
	TotalLayananChat      int `json:"total_layanan_chat"`
	TotalLayananVideoCall int `json:"total_layanan_video_call"`
}

type ManageResponse struct {
	DoctorID    uint   `json:"doctor_id"`
	PatientID   uint   `json:"patient_id"`
	PatientName string `json:"patient_name"`
	Gender      string `json:"gender"`
	Topic       string `json:"topic"`
	Layanan     string `json:"layanan"`
}

type DashboardAdminResponse struct {
	TotalDoctor        int `json:"total_doctor"`
	TotalDoctorBaru    int `json:"total_new_doctor"`
	TotalDoctorActive  int `json:"total_active_doctor"`
	TotalDoctorPending int `json:"total_pending_doctor"`
}
