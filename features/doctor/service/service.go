package service

import (
	"FinalProject/features/doctor"
	"FinalProject/helper"
	"FinalProject/helper/email"
	"FinalProject/utils/cloudinary"
	"errors"
	"math/rand"

	"github.com/sirupsen/logrus"
)

type DoctorService struct {
	data  doctor.DoctorDataInterface
	cld   cloudinary.CloudinaryInterface
	jwt   helper.JWTInterface
	email email.EmailInterface
}

func NewDoctor(data doctor.DoctorDataInterface, cloudinary cloudinary.CloudinaryInterface, email email.EmailInterface) doctor.DoctorServiceInterface {
	return &DoctorService{
		data:  data,
		cld:   cloudinary,
		email: email,
	}
}

func generateRandomCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

func (psvc *DoctorService) GetDoctors() ([]doctor.DoctorAll, error) {
	result, err := psvc.data.GetAll()
	if err != nil {
		return nil, errors.New("get All Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctor(id int) (*doctor.DoctorAll, error) {
	result, err := psvc.data.GetByID(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorByUserId(userID int) (*doctor.DoctorAll, error) {
	result, err := psvc.data.GetDoctorByUserId(userID)
	if err != nil {
		return nil, errors.New("get By User ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) SearchDoctor(name string) ([]doctor.DoctorAll, error) {
	result, err := psvc.data.SearchDoctor(name)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorWorkadays(id int) ([]doctor.DoctorWorkdays, error) {
	result, err := psvc.data.GetByIDWorkadays(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorEducation(id int) ([]doctor.DoctorEducation, error) {
	result, err := psvc.data.GetByIDEducation(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorExperience(id int) ([]doctor.DoctorExperience, error) {
	result, err := psvc.data.GetByIDExperience(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) CreateDoctor(newData doctor.Doctor) (*doctor.Doctor, error) {
	result, err := psvc.data.Insert(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}

	email, err := psvc.data.FindEmail(result.UserID)

	// emailUser, err := us.d.GetByEmail(email)

	// email := emailUser

	header := "Selamat " + result.DoctorName + ", pengajuan konselor Anda sudah kami terima!"
	htmlBody := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Pengajuan Konselor</title>
	</head>
	<body style="margin: 0; padding: 0; box-sizing: border-box;">
		<table align="center" cellpadding="0" cellspacing="0" width="95%">
		<tr>
			<td align="center">
			<table align="center" cellpadding="0" cellspacing="0" width="600" style="border-spacing: 2px 5px;" bgcolor="#fff">
				<tr>
				<td style="padding: 5px 5px 5px 5px;">
					<a href="#" target="_blank">
					<img src="https://i.ibb.co/kgMjHSV/Logo.png" alt="Logo" style="width:200px; border:0; margin:0;"/>
					</a>
				</td>
				</tr>
				<tr>
				<td bgcolor="#fff">
					<table cellpadding="0" cellspacing="0" width="100%%">
					<tr>
						<td style="padding: 10px 0 10px 0; font-family: Nunito, sans-serif; font-size: 18px; font-weight: 900">
						Halo, ` + result.DoctorName + `
						</td>
					</tr>
					</table>
				</td>
				</tr>
				<tr>
				<td bgcolor="#fff">
					<table cellpadding="0" cellspacing="0" width="100%%">
					<tr>
						<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
						Selamat! Kami dengan senang hati ingin memberitahu Anda bahwa pengajuan Anda sebagai Konselor di EmpathiCare telah berhasil diterima. Kami sangat berterima kasih atas ketertarikan Anda untuk bergabung dengan kami.
			<p></p>
						</td>
					</tr>
					<tr>
						<td style="padding: 0; font-family: Nunito, sans-serif; font-size: 16px;">
						Terima kasih atas kontribusi Anda dalam meningkatkan pelayanan kesehatan kami.
			 <p></p>
						</td>
					</tr>
		  <tr>
						<td style="padding: 10px 0 10px 0; font-family: Nunito, sans-serif; font-size: 16px; font-weight: 900">
						Salam Sehat,
						</td>
					</tr>
		   <tr>
						<td style="font-family: Nunito, sans-serif; font-size: 16px; font-weight: 900">
						Team EmpathiCare
						</td>
					</tr>
					</table>
				</td>
				</tr>
			</table>
			</td>
		</tr>
		</table>
	</body>
	</html>`

	ress := psvc.email.SendEmail(*email, header, htmlBody)

	logrus.Info("Info send email ==[]==", nil)
	logrus.Info("Email:", email, &email, *email)
	logrus.Info("UserID:", result.UserID)
	logrus.Info("Header: ", header)
	logrus.Info("HTML Body: ", htmlBody)
	logrus.Info("Result Pengiriman Email: ", ress)
	logrus.Info("Info send email ==[]==", nil)

	if ress != nil {
		return nil, ress
	}
	return result, nil
}

func (psvc *DoctorService) CreateDoctorExpertise(newData doctor.DoctorExpertiseRelation) (*doctor.DoctorExpertiseRelation, error) {
	resultExpertise, err := psvc.data.InsertExpertise(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return resultExpertise, nil
}

func (psvc *DoctorService) CreateDoctorWorkadays(newData doctor.DoctorWorkdays) (*doctor.DoctorWorkdays, error) {
	resultWorkadays, err := psvc.data.InsertWorkadays(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return resultWorkadays, nil
}

func (psvc *DoctorService) CreateDoctorEducation(newData doctor.DoctorEducation) (*doctor.DoctorEducation, error) {
	resultEducation, err := psvc.data.InsertEducation(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return resultEducation, nil
}

func (psvc *DoctorService) CreateDoctorExperience(newData doctor.DoctorExperience) (*doctor.DoctorExperience, error) {
	resultExperience, err := psvc.data.InsertExperience(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
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
		return false, errors.New("Update Datapokok Dokter Failed")
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
