package service

import (
	bundlecounseling "FinalProject/features/bundle_counseling"
	"FinalProject/utils/cloudinary"
	"errors"
)

type BundleCounselingService struct {
	d   bundlecounseling.BundleCounselingDataInterface
	cld cloudinary.CloudinaryInterface
}

func New(data bundlecounseling.BundleCounselingDataInterface, cld cloudinary.CloudinaryInterface) bundlecounseling.BundleCounselingServiceInterface {
	return &BundleCounselingService{
		d:   data,
		cld: cld,
	}
}

func (s *BundleCounselingService) GetAllBundle() ([]bundlecounseling.BundleCounselingInfo, error) {
	result, err := s.d.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *BundleCounselingService) CreateBundle(input bundlecounseling.BundleCounseling, file bundlecounseling.BundleCounselingFile) (*bundlecounseling.BundleCounseling, error) {

	uploadUrl, err := s.cld.UploadImageHelper(file.Avatar)
	if err != nil {
		return nil, errors.New("Upload Failed")
	}

	newData := bundlecounseling.BundleCounseling{
		Name:         input.Name,
		Sessions:     input.Sessions,
		Type:         input.Type,
		Price:        input.Price,
		Description:  input.Description,
		ActivePriode: input.ActivePriode,
		Avatar:       uploadUrl,
	}

	result, err := s.d.Create(newData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
