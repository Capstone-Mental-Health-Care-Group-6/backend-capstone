package data

import (
	"FinalProject/features/withdraw"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WithdrawData struct {
	db *gorm.DB
}

func New(db *gorm.DB) withdraw.WithdrawDataInterface {
	return &WithdrawData{
		db: db,
	}
}

func (ad *WithdrawData) GetAll() ([]withdraw.WithdrawInfo, error) {
	var list = []withdraw.WithdrawInfo{}
	var qry = ad.db.Table("withdraws").Select("withdraws.*, doctors.name as doctor_name, users.name as confirm_name").
		Joins("JOIN users on users.id = withdraws.confirm_by_id").
		Joins("JOIN doctors on doctors.id = withdraws.user_id").
		Where("withdraws.deleted_at is null").Scan(&list)

	if err := qry.Error; err != nil {
		logrus.Info("DB error : ", err.Error())
		return nil, err
	}

	return list, nil
}
