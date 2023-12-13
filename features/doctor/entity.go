package doctor

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Doctor struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	DoctorName string `json:"doctor_name"`

	DoctorNIK         string `json:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_numberphone"`
	DoctorGender      string `json:"doctor_gender"`
	DoctorAvatar      string `json:"doctor_avatar"`
	DoctorDescription string `json:"doctor_description"`

	DoctorMeetLink string `json:"doctor_meet_link"`
	DoctorSIPP     string `json:"doctor_sipp"`
	DoctorSIPPFile string `json:"doctor_sipp_file"`
	DoctorSTR      string `json:"doctor_str"`
	DoctorSTRFile  string `json:"doctor_str_file"`
	DoctorCV       string `json:"doctor_cv"`
	DoctorIjazah   string `json:"doctor_ijazah"`

	DoctorBalance uint   `json:"doctor_balance"`
	DoctorStatus  string `json:"doctor_status"`

	DoctorExpertise uint `json:"doctor_expertise"`

	DoctorWorkday    []DoctorWorkdays   `json:"workday"`
	DoctorEducation  []DoctorEducation  `json:"education"`
	DoctorExperience []DoctorExperience `json:"experience"`
	DoctorRating     []DoctorRating     `json:"ratings"`
}

type DoctorAll struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	DoctorName string `json:"doctor_name"`

	DoctorNIK         string `json:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_numberphone"`
	DoctorGender      string `json:"doctor_gender"`
	DoctorAvatar      string `json:"doctor_avatar"`
	DoctorDescription string `json:"doctor_description"`

	DoctorMeetLink string `json:"doctor_meet_link"`
	DoctorSIPP     string `json:"doctor_sipp"`
	DoctorSIPPFile string `json:"doctor_sipp_file"`
	DoctorSTR      string `json:"doctor_str"`
	DoctorSTRFile  string `json:"doctor_str_file"`
	DoctorCV       string `json:"doctor_cv"`
	DoctorIjazah   string `json:"doctor_ijazah"`

	DoctorBalance   uint   `json:"doctor_balance"`
	DoctorStatus    string `json:"doctor_status"`
	DoctorExpertise uint   `json:"doctor_expertise"`

	DoctorExperience []DoctorExperience `json:"experience" gorm:"foreignkey:DoctorID"`
	DoctorEducation  []DoctorEducation  `json:"education" gorm:"foreignkey:DoctorID"`
	DoctorWorkdays   []DoctorWorkdays   `json:"workday" gorm:"foreignkey:DoctorID"`
	DoctorRating     []DoctorRating     `json:"ratings" gorm:"foreignkey:DoctorID"`
}

type DoctorDatapokokUpdate struct {
	UserID            uint   `json:"user_id" form:"user_id"`
	DoctorName        string `json:"doctor_name" form:"doctor_name"`
	DoctorNIK         string `json:"doctor_nik"`
	DoctorDOB         string `json:"doctor_dob"`
	DoctorProvinsi    string `json:"doctor_provinsi"`
	DoctorKota        string `json:"doctor_kota"`
	DoctorNumberPhone string `json:"doctor_numberphone"`
	DoctorGender      string `json:"doctor_gender"`
	DoctorAvatar      string `json:"doctor_avatar"`
	DoctorDescription string `json:"doctor_description"`

	DoctorMeetLink string `json:"doctor_meet_link"`
	DoctorSIPP     string `json:"doctor_sipp"`
	DoctorSIPPFile string `json:"doctor_sipp_file"`
	DoctorSTR      string `json:"doctor_str"`
	DoctorSTRFile  string `json:"doctor_str_file"`
	DoctorCV       string `json:"doctor_cv"`
	DoctorIjazah   string `json:"doctor_ijazah"`
	DoctorBalance  uint   `json:"doctor_balance"`
	DoctorStatus   string `json:"doctor_status"`
	//FOR DOCTOR EXPERTISE
	DoctorExpertiseID uint `json:"expertise_id" form:"expertise_id"`
}

// type DoctorInfoWorkday struct {
// 	ID        uint      `json:"id"`
// 	DoctorID  uint      `json:"doctor_id"`
// 	WorkdayID uint      `json:"workday_id"`
// 	StartTime time.Time `json:"start_time"`
// 	EndTime   time.Time `json:"end_time"`
// }

// type DoctorInfoEducation struct {
// 	ID                 uint      `json:"id"`
// 	DoctorID           uint      `json:"doctor_id"`
// 	DoctorUniversity   string    `json:"doctor_university"`
// 	DoctorStudyProgram string    `json:"doctor_study_program"`
// 	DoctorEnrollYear   time.Time `json:"doctor_enroll_year"`
// 	DoctorGraduateYear time.Time `json:"doctor_graduate_year"`
// }

// type DoctorInfoExperience struct {
// 	ID                   uint      `json:"id"`
// 	DoctorID             uint      `json:"doctor_id"`
// 	DoctorCompany        string    `json:"doctor_company"`
// 	DoctorTitle          string    `json:"doctor_title"`
// 	DoctorCompanyAddress string    `json:"doctor_company_address"`
// 	DoctorStartDate      time.Time `json:"doctor_start_date"`
// 	DoctorEndDate        time.Time `json:"doctor_end_date"`
// }

type DoctorExpertiseRelation struct {
	DoctorID    uint `json:"doctor_id"`
	ExpertiseID uint `json:"expertise_id"`
}

type DoctorWorkdays struct {
	ID        uint      `json:"id"`
	DoctorID  uint      `json:"doctor_id"`
	WorkdayID uint      `json:"workday_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type DoctorRating struct {
	ID               uint   `json:"id"`
	DoctorID         uint   `json:"doctor_id"`
	PatientID        uint   `json:"patient_id"`
	TransactionID    string `json:"transaction_id"`
	DoctorStarRating uint   `json:"doctor_star_rating"`
	DoctorReview     string `json:"doctor_review"`
}

type DoctorEducation struct {
	ID                 uint      `json:"id"`
	DoctorID           uint      `json:"doctor_id"`
	DoctorUniversity   string    `json:"doctor_university"`
	DoctorStudyProgram string    `json:"doctor_study_program"`
	DoctorEnrollYear   time.Time `json:"doctor_enroll_year"`
	DoctorGraduateYear time.Time `json:"doctor_graduate_year"`
}

type DoctorExperience struct {
	ID                   uint      `json:"id"`
	DoctorID             uint      `json:"doctor_id"`
	DoctorCompany        string    `json:"doctor_company"`
	DoctorTitle          string    `json:"doctor_title"`
	DoctorCompanyAddress string    `json:"doctor_company_address"`
	DoctorStartDate      time.Time `json:"doctor_start_date"`
	DoctorEndDate        time.Time `json:"doctor_end_date"`
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

type DoctorHandlerInterface interface {
	GetDoctors() echo.HandlerFunc
	GetDoctor() echo.HandlerFunc
	GetDoctorByUserId() echo.HandlerFunc
	CreateDoctor() echo.HandlerFunc
	SearchDoctor() echo.HandlerFunc
	UpdateDoctorDatapokok() echo.HandlerFunc
	UpdateDoctorWorkdays() echo.HandlerFunc
	UpdateDoctorEducation() echo.HandlerFunc
	UpdateDoctorExperience() echo.HandlerFunc
	UpdateDoctorRating() echo.HandlerFunc
	InsertDataDoctor() echo.HandlerFunc
	DeleteDoctor() echo.HandlerFunc
	DeleteDoctorData() echo.HandlerFunc
	// DeleteWorkday() echo.HandlerFunc
	// DeleteEducation() echo.HandlerFunc
	// DeleteExperience() echo.HandlerFunc
}

type DoctorServiceInterface interface {
	GetDoctors() ([]DoctorAll, error)
	SearchDoctor(name string) ([]DoctorAll, error)
	GetDoctor(id int) (*DoctorAll, error)
	GetDoctorByUserId(userID int) (*DoctorAll, error)
	GetDoctorExperience(id int) ([]DoctorExperience, error)
	GetDoctorEducation(id int) ([]DoctorEducation, error)
	GetDoctorWorkadays(id int) ([]DoctorWorkdays, error)
	CreateDoctorExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	CreateDoctorWorkadays(newData DoctorWorkdays) (*DoctorWorkdays, error)
	CreateDoctorEducation(newData DoctorEducation) (*DoctorEducation, error)
	CreateDoctorExperience(newData DoctorExperience) (*DoctorExperience, error)
	CreateDoctor(newData Doctor) (*Doctor, error)
	DoctorAvatarUpload(newData DoctorAvatarPhoto) (string, error)
	DoctorSIPPUpload(newData DoctorSIPPFileDataModel) (string, error)
	DoctorSTRUpload(newData DoctorSTRFileDataModel) (string, error)
	DoctorCVUpload(newData DoctorCVDataModel) (string, error)
	DoctorIjazahUpload(newData DoctorIjazahDataModel) (string, error)
	UpdateDoctorDatapokok(id int, newData DoctorDatapokokUpdate) (bool, error)
	UpdateDoctorExperience(id int, doctorID int, newData DoctorExperience) (bool, error)
	UpdateDoctorWorkdays(id int, doctorID int, newData DoctorWorkdays) (bool, error)
	UpdateDoctorEducation(id int, doctorID int, newData DoctorEducation) (bool, error)
	UpdateDoctorRating(id int, patientID int, newData DoctorRating) (bool, error)
	DeleteDoctor(doctorID int) (bool, error)
	DeleteDoctorExperience(doctorID int) (bool, error)
	DeleteDoctorWorkdays(doctorID int) (bool, error)
	DeleteDoctorEducation(doctorID int) (bool, error)
	DeleteDoctorRating(doctorID int) (bool, error)
}

type DoctorDataInterface interface {
	GetAll() ([]DoctorAll, error)
	GetByID(id int) (*DoctorAll, error)
	GetDoctorByUserId(userID int) (*DoctorAll, error)
	GetByIDEducation(id int) ([]DoctorEducation, error)
	GetByIDWorkadays(id int) ([]DoctorWorkdays, error)
	GetByIDExperience(id int) ([]DoctorExperience, error)
	SearchDoctor(name string) ([]DoctorAll, error)
	Insert(newData Doctor) (*Doctor, error)
	InsertExpertise(newData DoctorExpertiseRelation) (*DoctorExpertiseRelation, error)
	InsertWorkadays(newData DoctorWorkdays) (*DoctorWorkdays, error)
	InsertEducation(newData DoctorEducation) (*DoctorEducation, error)
	InsertExperience(newData DoctorExperience) (*DoctorExperience, error)
	FindEmail(userID uint) (*string, error)
	UpdateDoctorDatapokok(id int, newData DoctorDatapokokUpdate) (bool, error)
	UpdateDoctorExperience(id int, doctorID int, newData DoctorExperience) (bool, error)
	UpdateDoctorWorkdays(id int, doctorID int, newData DoctorWorkdays) (bool, error)
	UpdateDoctorEducation(id int, doctorID int, newData DoctorEducation) (bool, error)
	UpdateDoctorRating(id int, pateintID int, newData DoctorRating) (bool, error)
	DeleteDoctor(doctorID int) (bool, error)
	// DeleteDoctorDatapokok(doctorID int) (bool, error)
	DeleteDoctorExperience(doctorID int) (bool, error)
	DeleteDoctorWorkdays(doctorID int) (bool, error)
	DeleteDoctorEducation(doctorID int) (bool, error)
	DeleteDoctorRating(doctorID int) (bool, error)
}
