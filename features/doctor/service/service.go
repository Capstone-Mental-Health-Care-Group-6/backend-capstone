package service

import (
	"FinalProject/features/doctor"
	"FinalProject/helper"
	"FinalProject/helper/email"
	"FinalProject/utils/cloudinary"
	"errors"

	"github.com/sirupsen/logrus"
)

type DoctorService struct {
	data  doctor.DoctorDataInterface
	cld   cloudinary.CloudinaryInterface
	email email.EmailInterface
	meet  helper.MeetInterface
}

func NewDoctor(data doctor.DoctorDataInterface, cloudinary cloudinary.CloudinaryInterface, email email.EmailInterface, meet helper.MeetInterface) doctor.DoctorServiceInterface {
	return &DoctorService{
		data:  data,
		cld:   cloudinary,
		email: email,
		meet:  meet,
	}
}

func (psvc *DoctorService) GetDoctors(name string) ([]doctor.DoctorAll, error) {
	result, err := psvc.data.GetAll(name)
	if err != nil {
		return nil, errors.New("Get All Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctor(id int) (*doctor.DoctorAll, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return nil, errors.New("Get By ID Doctor Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorByUserId(userID int) (*doctor.DoctorAll, error) {
	result, err := psvc.data.GetDoctorByUserId(userID)
	if err != nil {
		return nil, errors.New("Get Doctor By User ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorWorkadays(id int) ([]doctor.DoctorWorkdays, error) {
	result, err := psvc.data.GetByIDWorkadays(id)
	if err != nil {
		return nil, errors.New("Get By ID Workadays Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorEducation(id int) ([]doctor.DoctorEducation, error) {
	result, err := psvc.data.GetByIDEducation(id)
	if err != nil {
		return nil, errors.New("Get By ID Education Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorExperience(id int) ([]doctor.DoctorExperience, error) {
	result, err := psvc.data.GetByIDExperience(id)
	if err != nil {
		return nil, errors.New("Get By ID Experience Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) CreateDoctor(newData doctor.Doctor) (*doctor.Doctor, error) {
	result, err := psvc.data.Insert(newData)
	if err != nil {
		return nil, errors.New("Insert Doctor Process Failed")
	}

	email, err := psvc.data.FindEmail(result.UserID)
	if err != nil {
		return nil, errors.New("Find Email Process Failed")
	}

	header, htmlBody := psvc.email.HtmlBodyRegistDoctor(result.DoctorName)

	err = psvc.email.SendEmail(*email, header, htmlBody)

	logrus.Info("Info send email ==[]==", nil)
	logrus.Info("Email:", email, &email, *email)
	logrus.Info("UserID:", result.UserID)
	logrus.Info("Header: ", header)
	logrus.Info("HTML Body: ", htmlBody)
	logrus.Info("Result Pengiriman Email: ", err)
	logrus.Info("Info send email ==[]==", nil)

	if err != nil {
		return nil, errors.New("Send Email Error")
	}
	return result, nil
}

func (psvc *DoctorService) CreateDoctorExpertise(newData doctor.DoctorExpertiseRelation) (*doctor.DoctorExpertiseRelation, error) {
	resultExpertise, err := psvc.data.InsertExpertise(newData)
	if err != nil {
		return nil, errors.New("Insert Doctor Expertise Process Failed")
	}
	return resultExpertise, nil
}

func (psvc *DoctorService) CreateDoctorWorkadays(newData doctor.DoctorWorkdays) (*doctor.DoctorWorkdays, error) {
	resultWorkadays, err := psvc.data.InsertWorkadays(newData)
	if err != nil {
		return nil, errors.New("Insert Doctor Workadays Process Failed")
	}
	return resultWorkadays, nil
}

func (psvc *DoctorService) CreateDoctorEducation(newData doctor.DoctorEducation) (*doctor.DoctorEducation, error) {
	resultEducation, err := psvc.data.InsertEducation(newData)
	if err != nil {
		return nil, errors.New("Insert Doctor Education Process Failed")
	}
	return resultEducation, nil
}

func (psvc *DoctorService) CreateDoctorExperience(newData doctor.DoctorExperience) (*doctor.DoctorExperience, error) {
	resultExperience, err := psvc.data.InsertExperience(newData)
	if err != nil {
		return nil, errors.New("Insert Doctor Experience Process Failed")
	}
	return resultExperience, nil
}

func (psvc *DoctorService) DoctorAvatarUpload(newData doctor.DoctorAvatarPhoto) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorAvatar)
	if err != nil {
		return "", errors.New("Upload Avatar Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorSTRUpload(newData doctor.DoctorSTRFileDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorSTRFile)
	if err != nil {
		return "", errors.New("Upload Surat Tanda Registrasi Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorSIPPUpload(newData doctor.DoctorSIPPFileDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorSIPPFile)
	if err != nil {
		return "", errors.New("Upload SIPP Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorCVUpload(newData doctor.DoctorCVDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorCV)
	if err != nil {
		return "", errors.New("Upload CV Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) DoctorIjazahUpload(newData doctor.DoctorIjazahDataModel) (string, error) {
	uploadUrl, err := psvc.cld.UploadImageHelper(newData.DoctorIjazah)
	if err != nil {
		return "", errors.New("Upload Ijazah Failed")
	}
	return uploadUrl, nil
}

func (psvc *DoctorService) GetMeetLink() (string, error) {
	allMeetLinks := psvc.meet.GetMeetLink()

	for _, meetLink := range allMeetLinks {
		isLink := psvc.data.IsLinkUsed(meetLink)
		if !isLink {
			return meetLink, nil
		}
	}
	return "", errors.New("Semua link sudah digunakan")
}

func (psvc *DoctorService) UpdateDoctorDatapokok(id int, newData doctor.DoctorDatapokokUpdate) (bool, error) {
	result, err := psvc.data.UpdateDoctorDatapokok(id, newData)
	if err != nil {
		return false, errors.New("Update Datapokok Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) UpdateDoctorEducation(id int, doctorID int, newData doctor.DoctorEducation) (bool, error) {
	result, err := psvc.data.UpdateDoctorEducation(id, doctorID, newData)
	if err != nil {
		return false, errors.New("Update Education Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) UpdateDoctorExperience(id int, doctorID int, newData doctor.DoctorExperience) (bool, error) {
	result, err := psvc.data.UpdateDoctorExperience(id, doctorID, newData)
	if err != nil {
		return false, errors.New("Update Experience Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) UpdateDoctorWorkdays(id int, doctorID int, newData doctor.DoctorWorkdays) (bool, error) {
	result, err := psvc.data.UpdateDoctorWorkdays(id, doctorID, newData)
	if err != nil {
		return false, errors.New("Update Workdays Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) UpdateDoctorRating(id int, patientID int, newData doctor.DoctorRating) (bool, error) {
	result, err := psvc.data.UpdateDoctorRating(id, patientID, newData)
	if err != nil {
		return false, errors.New("Update Rating Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DeleteDoctor(doctorID int) (bool, error) {
	result, err := psvc.data.DeleteDoctor(doctorID)
	if err != nil {
		return false, errors.New("Delete Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DeleteDoctorEducation(doctorID int) (bool, error) {
	result, err := psvc.data.DeleteDoctorEducation(doctorID)
	if err != nil {
		return false, errors.New("Delete Education Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DeleteDoctorExperience(doctorID int) (bool, error) {
	result, err := psvc.data.DeleteDoctorExperience(doctorID)
	if err != nil {
		return false, errors.New("Delete Experience Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DeleteDoctorWorkdays(doctorID int) (bool, error) {
	result, err := psvc.data.DeleteDoctorWorkdays(doctorID)
	if err != nil {
		return false, errors.New("Delete Workdays Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DeleteDoctorRating(doctorID int) (bool, error) {
	result, err := psvc.data.DeleteDoctorRating(doctorID)
	if err != nil {
		return false, errors.New("Delete Rating Dokter Failed")
	}
	return result, nil
}

func (psvc *DoctorService) DoctorDashboard(id int) (doctor.DoctorDashboard, error) {
	res, err := psvc.data.DoctorDashboard(id)

	if err != nil {
		return res, errors.New("Process Failed")
	}

	return res, nil
}

func (psvc *DoctorService) DoctorDashboardPatient(id int) ([]doctor.DoctorDashboardPatient, error) {
	res, err := psvc.data.DoctorDashboardPatient(id)

	if err != nil {
		return res, errors.New("Process Failed")
	}

	return res, nil
}

func (psvc *DoctorService) DoctorDashboardAdmin() (doctor.DoctorDashboardAdmin, error) {
	res, err := psvc.data.DoctorDashboardAdmin()

	if err != nil {
		return res, errors.New("Process Failed")
	}

	return res, nil
}

func (psvc *DoctorService) DenyDoctor(userID int) (bool, error) {
	dataDoctor, err := psvc.data.GetDoctorByUserId(userID)
	if err != nil {
		return false, errors.New("Get Doctor By User ID Error")
	}

	doctorID := int(dataDoctor.ID)

	result, err := psvc.data.DenyDoctor(doctorID)

	if err != nil {
		return false, errors.New("Deny Process Failed")
	}

	return result, nil
}

func (psvc *DoctorService) ApproveDoctor(userID int) (bool, error) {
	dataDoctor, err := psvc.data.GetDoctorByUserId(userID)
	if err != nil {
		return false, errors.New("Get Doctor By User ID Error")
	}

	doctorID := int(dataDoctor.ID)

	result, err := psvc.data.ApproveDoctor(doctorID)

	if err != nil {
		return false, errors.New("Approve Process Failed")
	}

	return result, nil
}
