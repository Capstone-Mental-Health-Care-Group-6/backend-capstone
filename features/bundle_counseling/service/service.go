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

func (s *BundleCounselingService) GetBundle(id int) (*bundlecounseling.BundleCounseling, error) {
	result, err := s.d.GetById(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *BundleCounselingService) UpdateBundle(id int, input bundlecounseling.BundleCounseling, file bundlecounseling.BundleCounselingFile) (bool, error) {

	newData := bundlecounseling.BundleCounseling{
		Name:         input.Name,
		Sessions:     input.Sessions,
		Type:         input.Type,
		Price:        input.Price,
		Description:  input.Description,
		ActivePriode: input.ActivePriode,
	}

	if file.Avatar != nil {
		uploadUrl, err := s.cld.UploadImageHelper(file.Avatar)
		if err != nil {
			return false, errors.New("Upload Failed")
		}

		newData.Avatar = uploadUrl
	}

	result, err := s.d.Update(id, newData)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *BundleCounselingService) DeleteBundle(id int) (bool, error) {
	result, err := s.d.Delete(id)
	if err != nil {
		return false, err
	}

	return result, nil
}
