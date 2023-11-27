package data

import (
	bundlecounseling "FinalProject/features/bundle_counseling"

	"gorm.io/gorm"
)

type BundleCounselingData struct {
	db *gorm.DB
}

func New(db *gorm.DB) bundlecounseling.BundleCounselingDataInterface {
	return &BundleCounselingData{
		db: db,
	}
}

func (bc *BundleCounselingData) GetAll() ([]bundlecounseling.BundleCounselingInfo, error) {
	var listBundleCounseling = []bundlecounseling.BundleCounselingInfo{}

	var qry = bc.db.Table("bundle_counseling").
		Select("bundle_counseling.*").
		Where("bundle_counseling.deleted_at is null").
		Where("bundle_counseling.type = ?", "PREMIUM").
		Scan(&listBundleCounseling)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listBundleCounseling, nil
}

func (bc *BundleCounselingData) Create(input bundlecounseling.BundleCounseling) (*bundlecounseling.BundleCounselingInfo, error) {
	var newBundleCounseling = &bundlecounseling.BundleCounseling{
		Name:         input.Name,
		Sessions:     input.Sessions,
		Type:         input.Type,
		Price:        input.Price,
		Description:  input.Description,
		ActivePriode: input.ActivePriode,
		Avatar:       input.Avatar,
	}

	return &bundlecounseling.BundleCounselingInfo{
		Name:         newBundleCounseling.Name,
		Sessions:     newBundleCounseling.Sessions,
		Type:         newBundleCounseling.Type,
		Price:        newBundleCounseling.Price,
		Description:  newBundleCounseling.Description,
		ActivePriode: newBundleCounseling.ActivePriode,
		Avatar:       newBundleCounseling.Avatar,
	}, nil

}
