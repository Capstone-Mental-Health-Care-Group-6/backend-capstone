package data

import (
	"FinalProject/features/transaction"
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

func (ad *TransactionData) GetAll() ([]transaction.TransactionInfo, error) {
	var listTransactions []transaction.TransactionInfo
	var qry = ad.db.Find(&listTransactions) // Fetch all transactions data from the table

	if qry.Error != nil {
		return nil, qry.Error
	}

	return listTransactions, nil
}

func (ad *TransactionData) GetByID(id int) ([]transaction.TransactionInfo, error) {
	var transactionInfo transaction.TransactionInfo
	var qry = ad.db.Table("transactions").Where("id = ?", id).First(&transactionInfo)

	if qry.Error != nil {
		return nil, qry.Error
	}

	return []transaction.TransactionInfo{transactionInfo}, nil
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

	dbData.PaymentStatus = newData.PaymentStatus
	dbData.PaymentType = newData.PaymentType

	fmt.Println("Ive succeed create payment status", newData.PaymentStatus)

	if err := ad.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ad *TransactionData) Delete(id int) (bool, error) {
	var deleteData = new(Transaction)

	if err := ad.db.Where("id = ?", id).Delete(&deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}
