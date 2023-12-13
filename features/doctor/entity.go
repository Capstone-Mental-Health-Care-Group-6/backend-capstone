package doctor

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Doctor struct {
	ID                  uint   `json:"id"`
	UserID              uint   `json:"user_id"`
	DoctorName          string `json:"doctor_name"`
	DoctorExperienced   string `json:"doctor_experienced"`
	DoctorDescription   string `json:"doctor_description"`
	DoctorAvatar        string `json:"doctor_avatar"`
	DoctorOfficeName    string `json:"doctor_office_name"`
	DoctorOfficeAddress string `json:"doctor_office_address"`
	DoctorOfficeCity    string `json:"doctor_office_city"`
	DoctorMeetLink      string `json:"doctor_meet_link"`
	DoctorSIPP          uint   `json:"doctor_sipp"`
	DoctorSIPPFile      string `json:"doctor_sipp_file"`
	DoctorSTR           uint   `json:"doctor_str"`
	DoctorSTRFile       string `json:"doctor_str_file"`
	DoctorCV            string `json:"doctor_cv"`
	DoctorIjazah        string `json:"doctor_ijazah"`
	DoctorBalance       uint   `json:"doctor_balance"`
	DoctorStatus        string `json:"doctor_status"`

	DoctorExpertise uint `json:"doctor_expertise"`

	DoctorWorkday    []DoctorWorkadays  `json:"workday"`
	DoctorEducation  []DoctorEducation  `json:"education"`
	DoctorExperience []DoctorExperience `json:"experience"`
}

type DoctorAll struct {
	ID                  uint   `json:"id"`
	UserID              uint   `json:"user_id"`
	DoctorName          string `json:"doctor_name"`
	DoctorExperienced   string `json:"doctor_experienced"`
	DoctorDescription   string `json:"doctor_description"`
	DoctorAvatar        string `json:"doctor_avatar"`
	DoctorOfficeName    string `json:"doctor_office_name"`
	DoctorOfficeAddress string `json:"doctor_office_address"`
	DoctorOfficeCity    string `json:"doctor_office_city"`
	DoctorMeetLink      string `json:"doctor_meet_link"`
	DoctorSIPP          uint   `json:"doctor_sipp"`
	DoctorSIPPFile      string `json:"doctor_sipp_file"`
	DoctorSTR           uint   `json:"doctor_str"`
	DoctorSTRFile       string `json:"doctor_str_file"`
	DoctorCV            string `json:"doctor_cv"`
	DoctorIjazah        string `json:"doctor_ijazah"`
	DoctorBalance       uint   `json:"doctor_balance"`
	DoctorStatus        string `json:"doctor_status"`

	DoctorExpertise uint `json:"doctor_expertise"`

	DoctorExperience []DoctorInfoExperience `json:"experience" gorm:"foreignkey:DoctorID"`
	DoctorEducation  []DoctorInfoEducation  `json:"education" gorm:"foreignkey:DoctorID"`
	DoctorWorkday    []DoctorInfoWorkday    `json:"workday" gorm:"foreignkey:DoctorID"`
}

type DoctorInfo struct {
	UserID              uint   `json:"user_id"`
	DoctorName          string `json:"doctor_name"`
	DoctorExperienced   string `json:"doctor_experienced"`
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

	DoctorExpertise uint `json:"doctor_expertise"`
}

type DoctorInfoWorkday struct {
	// DoctorWorkday []DoctorWorkadays `json:"doctor_workday"`
	ID        uint      `json:"id"`
	DoctorID  uint      `json:"doctor_id"`
	WorkdayID uint      `json:"workday_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type DoctorInfoEducation struct {
	// DoctorEducation    []DoctorEducation `json:"doctor_education"`
	ID                 uint      `json:"id"`
	DoctorID           uint      `json:"doctor_id"`
	DoctorUniversity   string    `json:"doctor_university"`
	DoctorStudyProgram string    `json:"doctor_study_program"`
	DoctorGraduateYear time.Time `json:"doctor_graduate_year"`
}

type DoctorInfoExperience struct {
	// DoctorExperience []DoctorExperience `json:"doctor_experience"`
	ID                          uint      `json:"id"`
	DoctorID                    uint      `json:"doctor_id"`
	DoctorCompany               string    `json:"doctor_company"`
	DoctorTitle                 string    `json:"doctor_title"`
	DoctorExperienceDescription string    `json:"doctor_experience_description"`
	DoctorStartDate             time.Time `json:"doctor_start_date"`
	DoctorEndDate               time.Time `json:"doctor_end_date"`
	DoctorIsNow                 bool      `json:"doctor_is_now"`
}

// MODEL FOR REGISTERING

type DoctorExpertiseRelation struct {
	DoctorID    uint `json:"doctor_id"`
	ExpertiseID uint `json:"expertise_id"`
}

type DoctorWorkadays struct {
	DoctorID  uint      `json:"doctor_id"`
	WorkdayID uint      `json:"workday_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type DoctorRating struct {
	DoctorID         uint   `json:"doctor_id"`
	PatientID        uint   `json:"patient_id"`
	DoctorStarRating uint   `json:"doctor_star_rating"`
	DoctorReview     string `json:"doctor_review;type:varchar(255)"`
}

type DoctorEducation struct {
	DoctorID           uint      `json:"doctor_id"`
	DoctorUniversity   string    `json:"doctor_university"`
	DoctorStudyProgram string    `json:"doctor_study_program"`
	DoctorGraduateYear time.Time `json:"doctor_graduate_year"`
}

type DoctorExperience struct {
	DoctorID                    uint      `json:"doctor_id"`
	DoctorCompany               string    `json:"doctor_company"`
	DoctorTitle                 string    `json:"doctor_title"`
	DoctorExperienceDescription string    `json:"doctor_experience_description"`
	DoctorStartDate             time.Time `json:"doctor_start_date"`
	DoctorEndDate               time.Time `json:"doctor_end_date"`
	DoctorIsNow                 bool      `json:"doctor_is_now"`
}

type DoctorAvatarPhoto struct {
	DoctorAvatar multipart.File `json:"doctor_avatar"`
}

type DoctorSIPPFileDataModel struct {
	DoctorSIPPFile multipart.File `json:"doctor_sipp_file"`
}

type DoctorSTRFileDataModel struct {
	DoctorSTRFile multipart.File `json:"doctor_str_file"`
}

type DoctorCVDataModel struct {
	DoctorCV multipart.File `json:"doctor_cv"`
}

type DoctorIjazahDataModel struct {
	DoctorIjazah multipart.File `json:"doctor_ijazah"`
}

//MODEL FOR UPDATE

type DoctorDatapokokUpdate struct {
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

	//FOR DOCTOR EXPERTISE
	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id"`
}

type DoctorDashboard struct {
	TotalPatient          int `json:"total_patient"`
	TotalJamPraktek       int `json:"total_jam_praktek"`
	TotalLayananChat      int `json:"total_layanan_chat"`
	TotalLayananVideoCall int `json:"total_layanan_video_call"`
}

type DoctorDashboardPatient struct {
	PatientID   uint   `json:"patient_id"`
	PatientName string `json:"patient_name"`
	Gender      string `json:"gender"`
	Topic       string `json:"topic"`
	Layanan     string `json:"layanan"`
}

type DoctorManagePatient struct {
	DoctorID     uint   `json:"doctor_id"`
	PatientID    uint   `json:"patient_id"`
	Alasan       string `json:"alasan"`
	DetailAlasan string `json:"detail_alasan"`
}

type DoctorHandlerInterface interface {
	GetDoctors() echo.HandlerFunc
	GetDoctor() echo.HandlerFunc
	CreateDoctor() echo.HandlerFunc
	SearchDoctor() echo.HandlerFunc
	UpdateDoctorDatapokok() echo.HandlerFunc
	UpdateDoctorWorkdays() echo.HandlerFunc
	UpdateDoctorEducation() echo.HandlerFunc
	UpdateDoctorExperience() echo.HandlerFunc
	InsertWorkday() echo.HandlerFunc
	InsertEducation() echo.HandlerFunc
	InsertExperience() echo.HandlerFunc
	DeleteDoctor() echo.HandlerFunc
	DeleteWorkday() echo.HandlerFunc
	DeleteEducation() echo.HandlerFunc
	DeleteExperience() echo.HandlerFunc
	DoctorDashboard() echo.HandlerFunc
	DoctorDashboardPatient() echo.HandlerFunc
}

type DoctorServiceInterface interface {
	GetDoctors() ([]DoctorAll, error)
	SearchDoctor(name string) ([]DoctorAll, error)
	GetDoctor(id int) (*DoctorAll, error)
	GetDoctorExperience(id int) ([]DoctorInfoExperience, error)
	GetDoctorEducation(id int) ([]DoctorInfoEducation, error)
	GetDoctorWorkadays(id int) ([]DoctorInfoWorkday, error)
	CreateDoctorExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	CreateDoctorWorkadays(newData DoctorWorkadays) (*DoctorWorkadays, error)
	CreateDoctorEducation(newData DoctorEducation) (*DoctorEducation, error)
	CreateDoctorExperience(newData DoctorExperience) (*DoctorExperience, error)
	CreateDoctor(newData Doctor) (*Doctor, error)
	DoctorAvatarUpload(newData DoctorAvatarPhoto) (string, error)
	DoctorSIPPUpload(newData DoctorSIPPFileDataModel) (string, error)
	DoctorSTRUpload(newData DoctorSTRFileDataModel) (string, error)
	DoctorCVUpload(newData DoctorCVDataModel) (string, error)
	DoctorIjazahUpload(newData DoctorIjazahDataModel) (string, error)
	UpdateDoctorDatapokok(id int, newData DoctorDatapokokUpdate) (bool, error)
	UpdateDoctorExperience(id int, doctorID int, newData DoctorInfoExperience) (bool, error)
	UpdateDoctorWorkdays(id int, doctorID int, newData DoctorInfoWorkday) (bool, error)
	UpdateDoctorEducation(id int, doctorID int, newData DoctorInfoEducation) (bool, error)
	DoctorDashboard(id int) (DoctorDashboard, error)
	DoctorDashboardPatient(id int) ([]DoctorDashboardPatient, error)
}

type DoctorDataInterface interface {
	GetAll() ([]DoctorAll, error)
	GetByID(id int) (*DoctorAll, error)
	GetByIDEducation(id int) ([]DoctorInfoEducation, error)
	GetByIDWorkadays(id int) ([]DoctorInfoWorkday, error)
	GetByIDExperience(id int) ([]DoctorInfoExperience, error)
	SearchDoctor(name string) ([]DoctorAll, error)
	Insert(newData Doctor) (*Doctor, error)
	InsertExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	InsertWorkadays(newData DoctorWorkadays) (*DoctorWorkadays, error)
	InsertEducation(newData DoctorEducation) (*DoctorEducation, error)
	InsertExperience(newData DoctorExperience) (*DoctorExperience, error)
	FindEmail(userID uint) (*string, error)
	UpdateDoctorDatapokok(id int, newData DoctorDatapokokUpdate) (bool, error)
	UpdateDoctorExperience(id int, doctorID int, newData DoctorInfoExperience) (bool, error)
	UpdateDoctorWorkdays(id int, doctorID int, newData DoctorInfoWorkday) (bool, error)
	UpdateDoctorEducation(id int, doctorID int, newData DoctorInfoEducation) (bool, error)
	DoctorDashboard(id int) (DoctorDashboard, error)
	DoctorDashboardPatient(id int) ([]DoctorDashboardPatient, error)
	DoctorManagePatient(id int) (DoctorManagePatient, error)
}
