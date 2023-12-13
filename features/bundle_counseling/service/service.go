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
		return nil, errors.New("Get All Process Failed")
	}
	return result, nil
}

func (s *BundleCounselingService) GetAllBundleFilter(jenis string, metode int, durasi int) ([]bundlecounseling.BundleCounselingInfo, error) {

	var listBundleCounseling = []bundlecounseling.BundleCounselingInfo{}

	if jenis == "INSTAN" {
		metode = 1
		durasi = 1
	}

	hargaMetode, err := s.d.HargaMetode(metode)
	if err != nil {
		return listBundleCounseling, errors.New("Get Harga Metode Failed")
	}

	hargaDurasi, err := s.d.HargaDurasi(durasi)
	if err != nil {
		return listBundleCounseling, errors.New("Get Harga Durasi Failed")
	}

	bundles, err := s.d.GetAllFilter(jenis)
	if err != nil {
		return listBundleCounseling, errors.New("Get All Filter Process Failed")
	}

	for _, bundle := range bundles {
		listBundleCounseling = append(listBundleCounseling, bundlecounseling.BundleCounselingInfo{
			ID:           bundle.ID,
			Name:         bundle.Name,
			Sessions:     bundle.Sessions,
			Type:         bundle.Type,
			Price:        bundle.Price + hargaMetode + hargaDurasi,
			Description:  bundle.Description,
			ActivePriode: bundle.ActivePriode,
			Avatar:       bundle.Avatar,
		})
	}

	return listBundleCounseling, nil
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
		return nil, errors.New("Create Process Failed")
	}

	return result, nil
}

func (s *BundleCounselingService) GetBundle(id int) (*bundlecounseling.BundleCounseling, error) {
	result, err := s.d.GetById(id)
	if err != nil {
		return nil, errors.New("Get By ID Process Failed")
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
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (s *BundleCounselingService) DeleteBundle(id int) (bool, error) {
	result, err := s.d.Delete(id)
	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}
