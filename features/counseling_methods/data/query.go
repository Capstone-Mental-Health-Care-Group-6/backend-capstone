package data

import (
	counselingmethod "FinalProject/features/counseling_methods"

	"gorm.io/gorm"
)

type CounselingMethodData struct {
	db *gorm.DB
}

func New(db *gorm.DB) counselingmethod.CounselingMethodDataInterface {
	return &CounselingMethodData{
		db: db,
	}
}

func (cmd *CounselingMethodData) GetAll() ([]counselingmethod.CounselingMethodInfo, error) {
	var data = []counselingmethod.CounselingMethodInfo{}
	if err := cmd.db.Table("counseling_methods").Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (cmd *CounselingMethodData) GetByID(id int) ([]counselingmethod.CounselingMethodInfo, error) {
	var data = []counselingmethod.CounselingMethodInfo{}
	if err := cmd.db.Table("counseling_methods").Where("id = ?", id).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
