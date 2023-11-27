package handler

type DoctorResponse struct {
	ID                  uint   `json:"id"`
	UserID              uint   `json:"user_id"`
	DoctorName          string `json:"doctor_name"`
	DoctorExperience    string `json:"doctor_experience"`
	DoctorDescription   string `json:"doctor_description"`
	DoctorAvatar        string `json:"doctor_avatar"`
	DoctorOfficeName    string `json:"doctor_office_name"`
	DoctorOfficeAddress string `json:"doctor_office_address"`
	DoctorOfficeCity    string `json:"doctor_office_city"`
	DoctorMeetLink      string `json:"doctor_meet_link"`
	DoctorSIPPFile      string `json:"doctor_sipp_file"`
	DoctorSTRFile       string `json:"doctor_str_file"`
	DoctorCV            string `json:"doctor_cv"`
	DoctorIjazah        string `json:"doctor_ijazah"`
	DoctorBalance       uint   `json:"doctor_balance"`
	DoctorStatus        string `json:"doctor_status"`
}
