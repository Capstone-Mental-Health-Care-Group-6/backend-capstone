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
	DoctorID  uint      `json:"doctor_id"`
	WorkdayID uint      `json:"workday_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type DoctorInfoEducation struct {
	// DoctorEducation    []DoctorEducation `json:"doctor_education"`
	DoctorID           uint      `json:"doctor_id"`
	DoctorUniversity   string    `json:"doctor_university"`
	DoctorStudyProgram string    `json:"doctor_study_program"`
	DoctorGraduateYear time.Time `json:"doctor_graduate_year"`
}

type DoctorInfoExperience struct {
	// DoctorExperience []DoctorExperience `json:"doctor_experience"`
	DoctorID                    uint      `json:"doctor_id"`
	DoctorCompany               string    `json:"doctor_company"`
	DoctorTitle                 string    `json:"doctor_title"`
	DoctorExperienceDescription string    `json:"doctor_experience_description"`
	DoctorStartDate             time.Time `json:"doctor_start_date"`
	DoctorEndDate               time.Time `json:"doctor_end_date"`
	DoctorIsNow                 bool      `json:"doctor_is_now"`
}

// DoctorWorkadays string `json:"workday_id"`

// StartTime time.Time `json:"start_time"`
// EndTime   time.Time `json:"end_time"`

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

type JwtMapClaims struct {
	ID     uint `json:"id"`
	Role   uint `json:"role"`
	Status uint `json:"status"`
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

type DoctorHandlerInterface interface {
	GetDoctors() echo.HandlerFunc
	GetDoctor() echo.HandlerFunc
	CreateDoctor() echo.HandlerFunc
}

type DoctorServiceInterface interface {
	GetDoctors() ([]DoctorAll, error)
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
	JwtExtractToken(authorizationHeader string) (JwtMapClaims, error)
}

type DoctorDataInterface interface {
	GetAll() ([]DoctorAll, error)
	GetByID(id int) (*DoctorAll, error)
	GetByIDEducation(id int) ([]DoctorInfoEducation, error)
	GetByIDWorkadays(id int) ([]DoctorInfoWorkday, error)
	GetByIDExperience(id int) ([]DoctorInfoExperience, error)
	Insert(newData Doctor) (*Doctor, error)
	InsertExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	InsertWorkadays(newData DoctorWorkadays) (*DoctorWorkadays, error)
	InsertEducation(newData DoctorEducation) (*DoctorEducation, error)
	InsertExperience(newData DoctorExperience) (*DoctorExperience, error)
}
