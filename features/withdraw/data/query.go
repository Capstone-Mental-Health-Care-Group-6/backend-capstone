package data

import (
	"FinalProject/features/withdraw"
	"errors"

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
	var qry = ad.db.Table("withdraws").
		Select("withdraws.id, doctors.doctor_name as doctor_name, users.name as confirm_name, withdraws.balance_before, withdraws.balance_after, withdraws.balance_req, withdraws.payment_method, withdraws.payment_number, withdraws.payment_name, withdraws.date_confirmed, withdraws.status").
		Joins("LEFT JOIN doctors on doctors.id = withdraws.doctor_id").
		Joins("LEFT JOIN users on users.id = withdraws.confirm_by_id").
		Where("withdraws.deleted_at is null").Scan(&list)

	if err := qry.Error; err != nil {
		logrus.Info("DB error : ", err.Error())
		return nil, err
	}

	return list, nil
}

func (ad *WithdrawData) GetAllDoctor(id uint) ([]withdraw.WithdrawInfo, error) {
	var list = []withdraw.WithdrawInfo{}
	var qry = ad.db.Table("withdraws").
		Select("withdraws.id, doctors.doctor_name as doctor_name, users.name as confirm_name, withdraws.balance_before, withdraws.balance_after, withdraws.balance_req, withdraws.payment_method, withdraws.payment_number, withdraws.payment_name, withdraws.date_confirmed, withdraws.status").
		Joins("LEFT JOIN doctors on doctors.id = withdraws.doctor_id").
		Joins("LEFT JOIN users on users.id = withdraws.confirm_by_id").
		Where("withdraws.deleted_at is null").
		Where("withdraws.doctor_id = ?", id).Scan(&list)

	if err := qry.Error; err != nil {
		logrus.Info("DB error : ", err.Error())
		return nil, err
	}

	return list, nil
}

func (ad *WithdrawData) GetUserDoctor(id uint) (uint, error) {
	var doctorID uint
	var qry = ad.db.Table("users").
		Select("doctors.id as doctor_id").
		Joins("LEFT JOIN doctors on doctors.user_id = users.id").
		Where("users.id = ?", id).Scan(&doctorID)

	if err := qry.Error; err != nil {
		return 0, err
	}

	return doctorID, nil
}

func (ad *WithdrawData) Insert(newData withdraw.Withdraw) (*withdraw.Withdraw, error) {
	var data = new(Withdraw)
	data.DoctorID = newData.DoctorID
	data.BalanceBefore = newData.BalanceBefore
	data.BalanceAfter = newData.BalanceAfter
	data.BalanceReq = newData.BalanceReq
	data.PaymentMethod = newData.PaymentMethod
	data.PaymentNumber = newData.PaymentNumber
	data.PaymentName = newData.PaymentName

	if err := ad.db.Create(data).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *WithdrawData) GetBalance(idDoctor uint) (uint, error) {
	var balance uint
	var qry = ad.db.Table("doctors").
		Select("doctor_balance").
		Where("id = ?", idDoctor).Scan(&balance)

	if err := qry.Error; err != nil {
		return 0, err
	}

	return balance, nil
}

func (ad *WithdrawData) LessBalance(idDoctor uint, balance uint) (bool, error) {
	var qry = ad.db.Table("doctors").Where("id = ?", idDoctor).
		Update("doctor_balance", gorm.Expr("doctor_balance - ?", balance))

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (ad *WithdrawData) GetByID(id int) (*withdraw.WithdrawInfo, error) {
	var withdraw = new(withdraw.WithdrawInfo)
	var qry = ad.db.Table("withdraws").
		Select("withdraws.id, doctors.doctor_name as doctor_name, users.name as confirm_name, withdraws.balance_before, withdraws.balance_after, withdraws.balance_req, withdraws.payment_method, withdraws.payment_number, withdraws.payment_name, withdraws.date_confirmed, withdraws.status").
		Joins("LEFT JOIN doctors on doctors.id = withdraws.doctor_id").
		Joins("LEFT JOIN users on users.id = withdraws.confirm_by_id").
		Where("withdraws.deleted_at is null").
		Where("withdraws.id = ?", id).
		Scan(&withdraw)

	if err := qry.Error; err != nil {
		return nil, err
	}

	return withdraw, nil
}

func (ad *WithdrawData) UpdateStatus(id int, newData withdraw.Withdraw) (bool, error) {
	var qry = ad.db.Table("withdraws").
		Where("id = ?", id).
		Updates(Withdraw{
			Status:        newData.Status,
			ConfirmByID:   newData.ConfirmByID,
			DateConfirmed: newData.DateConfirmed,
		})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
