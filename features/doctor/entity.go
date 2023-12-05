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
	DoctorExperience    string `json:"doctor_experience"`
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

	DoctorWorkdayID     uint   `json:"workday_id"`
	DoctorExpertiseID   uint   `json:"expertise_id"`
	DoctorWorkStartTime string `json:"start_time"`
	DoctorWorkEndTime   string `json:"end_time"`
}

type DoctorInfo struct {
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

	WorkdayID   uint      `json:"workday_id"`
	ExpertiseID uint      `json:"expertise_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

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
	GetDoctors() ([]DoctorInfo, error)
	GetDoctor(id int) ([]DoctorInfo, error)
	CreateDoctorExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	CreateDoctorWorkadays(newData DoctorWorkadays) (*DoctorWorkadays, error)
	CreateDoctor(newData Doctor) (*Doctor, error)
	DoctorAvatarUpload(newData DoctorAvatarPhoto) (string, error)
	DoctorSIPPUpload(newData DoctorSIPPFileDataModel) (string, error)
	DoctorSTRUpload(newData DoctorSTRFileDataModel) (string, error)
	DoctorCVUpload(newData DoctorCVDataModel) (string, error)
	DoctorIjazahUpload(newData DoctorIjazahDataModel) (string, error)
	JwtExtractToken(authorizationHeader string) (JwtMapClaims, error)
}

type DoctorDataInterface interface {
	GetAll() ([]DoctorInfo, error)
	GetByID(id int) ([]DoctorInfo, error)
	Insert(newData Doctor) (*Doctor, error)
	InsertExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	InsertWorkadays(newData DoctorWorkadays) (*DoctorWorkadays, error)
}
