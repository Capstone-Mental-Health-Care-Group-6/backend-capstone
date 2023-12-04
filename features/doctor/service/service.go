package service

import (
	"FinalProject/features/doctor"
	"FinalProject/helper"
	"FinalProject/utils/cloudinary"
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type DoctorService struct {
	data doctor.DoctorDataInterface
	cld  cloudinary.CloudinaryInterface
	jwt  helper.JWTInterface
}

func NewDoctor(data doctor.DoctorDataInterface, cloudinary cloudinary.CloudinaryInterface) doctor.DoctorServiceInterface {
	return &DoctorService{
		data: data,
		cld:  cloudinary,
	}
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

func (psvc *DoctorService) GetDoctorWorkadays(id int) ([]doctor.DoctorInfoWorkday, error) {
	result, err := psvc.data.GetByIDWorkadays(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorEducation(id int) ([]doctor.DoctorInfoEducation, error) {
	result, err := psvc.data.GetByIDEducation(id)
	if err != nil {
		return nil, errors.New("get By ID Process Failed")
	}
	return result, nil
}

func (psvc *DoctorService) GetDoctorExperience(id int) ([]doctor.DoctorInfoExperience, error) {
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
	return result, nil
}

func (psvc *DoctorService) CreateDoctorExpertise(newData doctor.DoctorExpertiseRelation) (*doctor.DoctorExpertiseRelation, error) {
	resultExpertise, err := psvc.data.InsertExpertise(newData)
	if err != nil {
		return nil, errors.New("insert Process Failed")
	}
	return resultExpertise, nil
}

func (psvc *DoctorService) CreateDoctorWorkadays(newData doctor.DoctorWorkadays) (*doctor.DoctorWorkadays, error) {
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

func (psvc *DoctorService) JwtExtractToken(authorizationHeader string) (doctor.JwtMapClaims, error) {
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("SECRET"), nil
	})

	if err != nil {
		return doctor.JwtMapClaims{}, err
	}

	if token.Valid {
		var claims = token.Claims.(jwt.MapClaims)
		result := doctor.JwtMapClaims{}
		result.ID = claims["id"].(uint)
		result.Role = claims["role"].(uint)
		result.Status = claims["status"].(uint)
		return result, nil
	}

	return doctor.JwtMapClaims{}, nil
}
