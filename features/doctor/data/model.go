package data

import "gorm.io/gorm"

type Doctor struct {
	*gorm.Model
	UserID              uint   `gorm:"column:user_id"`
	DoctorName          string `gorm:"column:doctor_name;type:varchar(255)"`
	DoctorExperience    string `gorm:"column:doctor_experience;type:varchar(255)"`
	DoctorDescription   string `gorm:"column:doctor_description;type:varchar(255)"`
	DoctorAvatar        string `gorm:"column:doctor_avatar;type:varchar(255)"`
	DoctorOfficeName    string `gorm:"column:doctor_office_name;type:varchar(255)"`
	DoctorOfficeAddress string `gorm:"column:doctor_office_address;type:varchar(255)"`
	DoctorOfficeCity    string `gorm:"column:doctor_office_city;type:varchar(255)"`
	DoctorMeetLink      string `gorm:"column:doctor_meet_link;type:varchar(255)"`
	DoctorSIPP          uint   `gorm:"column:doctor_sipp"`
	DoctorSIPPFile      string `gorm:"column:doctor_sipp_file;type:varchar(255)"`
	DoctorSTR           uint   `gorm:"column:doctor_str"`
	DoctorSTRFile       string `gorm:"column:doctor_str_file;type:varchar(255)"`
	DoctorCV            string `gorm:"column:doctor_cv;type:varchar(255)"`
	DoctorIjazah        string `gorm:"column:doctor_ijazah;type:varchar(255)"`
	DoctorBalance       uint   `gorm:"column:doctor_balance"`
	DoctorStatus        string `gorm:"column:doctor_status;type:enum('active','not_active')"`
}
