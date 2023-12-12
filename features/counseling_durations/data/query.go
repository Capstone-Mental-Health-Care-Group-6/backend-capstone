package data

import (
	counselingduration "FinalProject/features/counseling_durations"

	"gorm.io/gorm"
)

type CounselingDurationData struct {
	db *gorm.DB
}

func New(db *gorm.DB) counselingduration.CounselingDurationDataInterface {
	return &CounselingDurationData{
		db: db,
	}
}

func (cmd *CounselingDurationData) GetAll() ([]counselingduration.CounselingDurationInfo, error) {
	var data = []counselingduration.CounselingDurationInfo{}
	if err := cmd.db.Table("counseling_durations").Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (cmd *CounselingDurationData) GetByID(id int) ([]counselingduration.CounselingDurationInfo, error) {
	var data = []counselingduration.CounselingDurationInfo{}
	if err := cmd.db.Table("counseling_durations").Where("id = ?", id).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
