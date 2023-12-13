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

func (bc *BundleCounselingData) GetAllFilter(jenis string) ([]bundlecounseling.BundleCounselingInfo, error) {
	var listBundleCounseling = []bundlecounseling.BundleCounselingInfo{}

	var qry = bc.db.Table("bundle_counseling").
		Select("bundle_counseling.*").
		Where("bundle_counseling.deleted_at is null").
		Where("bundle_counseling.type = ?", jenis).
		Scan(&listBundleCounseling)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return listBundleCounseling, nil
}

func (bc *BundleCounselingData) Create(input bundlecounseling.BundleCounseling) (*bundlecounseling.BundleCounseling, error) {
	// var newBundleCounseling = &bundlecounseling.BundleCounseling{
	// 	Name:         input.Name,
	// 	Sessions:     input.Sessions,
	// 	Type:         input.Type,
	// 	Price:        input.Price,
	// 	Description:  input.Description,
	// 	ActivePriode: input.ActivePriode,
	// 	Avatar:       input.Avatar,
	// }

	var newData = new(bundlecounseling.BundleCounseling)
	newData.Name = input.Name
	newData.Sessions = input.Sessions
	newData.Type = input.Type
	newData.Price = input.Price
	newData.Description = input.Description
	newData.ActivePriode = input.ActivePriode
	newData.Avatar = input.Avatar

	if err := bc.db.Table("bundle_counseling").Create(newData).Error; err != nil {
		return nil, err
	}

	return &input, nil
}

func (bc *BundleCounselingData) GetById(id int) (*bundlecounseling.BundleCounseling, error) {
	var result = new(bundlecounseling.BundleCounseling)

	var qry = bc.db.Table("bundle_counseling").
		Select("bundle_counseling.*").
		Where("bundle_counseling.deleted_at is null").
		Where("bundle_counseling.id = ?", id).
		Scan(&result)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (bc *BundleCounselingData) Update(id int, input bundlecounseling.BundleCounseling) (bool, error) {
	var newData = map[string]interface{}{
		"name":          input.Name,
		"sessions":      input.Sessions,
		"type":          input.Type,
		"price":         input.Price,
		"description":   input.Description,
		"active_priode": input.ActivePriode,
	}

	if input.Avatar != "" {
		newData["avatar"] = input.Avatar
	}

	if err := bc.db.Table("bundle_counseling").Where("id = ?", id).Updates(newData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (bc *BundleCounselingData) Delete(id int) (bool, error) {
	var deleteData = new(bundlecounseling.BundleCounseling)

	if err := bc.db.Table("bundle_counseling").Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (bc *BundleCounselingData) HargaMetode(id int) (uint, error) {
	var harga uint
	if err := bc.db.Table("counseling_methods").Select("additional_price").Where("id = ?", id).Scan(&harga).Error; err != nil {
		return 0, err
	}
	return harga, nil
}

func (bc *BundleCounselingData) HargaDurasi(id int) (uint, error) {
	var harga uint
	if err := bc.db.Table("counseling_durations").Select("additional_price").Where("id = ?", id).Scan(&harga).Error; err != nil {
		return 0, err
	}
	return harga, nil
}
