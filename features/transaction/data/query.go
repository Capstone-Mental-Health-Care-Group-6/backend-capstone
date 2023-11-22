package data

import (
	"FinalProject/features/transaction"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TransactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.TransactionDataInterface {
	return &TransactionData{
		db: db,
	}
}

func (ad *TransactionData) GetAndUpdate(newData transaction.UpdateTransaction, id string) (bool, error) {

	var transaction Transaction
	db := ad.db
	db.Where("midtrans_id = ?", id).First(&transaction)
	fmt.Println("This is the id: ", transaction.ID)
	transactionID := transaction.ID

	fmt.Println("This is the new payment status: ", newData.PaymentStatus)

	qry := db.Table("transactions").Where("id = ?", transactionID).Updates(Transaction{
		PaymentStatus: newData.PaymentStatus,
	})

	if qry.Error != nil {
		return false, nil
	}

	if err := qry.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (ad *TransactionData) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
	var transactionInfos []transaction.TransactionInfo
	var qry = ad.db.Table("transactions").Where("midtrans_id = ?", id).Select("user_id, midtrans_id, payment_status, payment_type, price_result").Find(&transactionInfos)

	if qry.Error != nil {
		return nil, qry.Error
	}

	fmt.Println("Json Response for query:", transactionInfos)

	return transactionInfos, nil
}

// func (ad *TransactionData) GetAll() ([]transaction.TransactionInfo, error) {
// 	var listTransactions []transaction.TransactionInfo            // Change to a slice to hold multiple transactions
// 	var qry = ad.db.Table("transactions").Find(&listTransactions) // Fetch all transactions data from the table

// 	if qry.Error != nil {
// 		return nil, qry.Error
// 	}

// 	return listTransactions, nil
// }

func (ad *TransactionData) GetAll(sort string) ([]transaction.TransactionInfo, error) {
	var listTransactions []transaction.TransactionInfo
	var qry = ad.db.Table("transactions")

	if sort != "" {
		qry = qry.Where("payment_type = ?", sort)
	}

	qry = qry.Find(&listTransactions)

	if qry.Error != nil {
		return nil, qry.Error
	}

	return listTransactions, nil
}

func (ad *TransactionData) GetByID(id int, sort string) ([]transaction.Transaction, error) {
	var transactionInfo []transaction.Transaction
	var qry = ad.db.Table("transactions").Where("user_id = ?", id)

	if sort != "" {
		qry = qry.Where("payment_type = ?", sort)
	}

	qry = qry.Find(&transactionInfo)

	if qry.Error != nil {
		return nil, qry.Error
	}

	return transactionInfo, nil
}

func (ad *TransactionData) Insert(newData transaction.Transaction) (*transaction.Transaction, error) {
	var dbData = new(Transaction)

	dbData.TopicID = newData.TopicID
	dbData.PatientID = newData.PatientID
	dbData.DoctorID = newData.DoctorID
	dbData.MethodID = newData.MethodID
	dbData.DurationID = newData.DurationID
	dbData.CounselingID = newData.CounselingID
	dbData.UserID = newData.UserID
	dbData.MidtransID = newData.MidtransID

	dbData.CounselingSession = newData.CounselingSession
	dbData.CounselingType = newData.CounselingType

	dbData.PriceMethod = newData.PriceMethod
	dbData.PriceDuration = newData.PriceDuration
	dbData.PriceCounseling = newData.PriceCounseling
	dbData.PriceResult = newData.PriceResult
	dbData.PaymentProof = newData.PaymentProof

	dbData.PaymentStatus = newData.PaymentStatus
	dbData.PaymentType = newData.PaymentType

	fmt.Println("Ive succeed create payment status", newData.PaymentStatus)

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *TransactionData) Delete(id int) (bool, error) {
	var transactionInfo transaction.Transaction
	var qry = ad.db.Table("transactions").Where("id = ?", id).Delete(&transactionInfo)

	if qry.Error != nil {
		return false, qry.Error
	}

	return true, nil
}

func (ad *TransactionData) Update(newData transaction.UpdateTransactionManual, id int) (bool, error) {
	var qry = ad.db.Table("transactions").Where("id = ?", id).Updates(Transaction{
		UserID:          newData.UserID,
		PriceMethod:     newData.PriceMethod,
		PriceDuration:   newData.PriceDuration,
		PriceCounseling: newData.PriceCounseling,
		PriceResult:     newData.PriceResult,
		PaymentStatus:   newData.PaymentStatus,
		PaymentType:     newData.PaymentType})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}

func (ad *TransactionData) UpdateWithTrxID(newData transaction.UpdateTransactionManual, id string) (bool, error) {
	var qry = ad.db.Table("transactions").Where("midtrans_id = ?", id).Updates(Transaction{
		UserID:          newData.UserID,
		PriceMethod:     newData.PriceMethod,
		PriceDuration:   newData.PriceDuration,
		PriceCounseling: newData.PriceCounseling,
		PriceResult:     newData.PriceResult,
		PaymentStatus:   newData.PaymentStatus,
		PaymentType:     newData.PaymentType})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("Update Data Error, No Data Affected")
	}

	return true, nil
}
