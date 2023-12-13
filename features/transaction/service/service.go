package service

import (
	transaction "FinalProject/features/transaction"
	"FinalProject/utils/cloudinary"
	"FinalProject/utils/midtrans"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type TransactionService struct {
	d   transaction.TransactionDataInterface
	cld cloudinary.CloudinaryInterface
	mt  midtrans.MidtransService
}

func New(data transaction.TransactionDataInterface, cloudinary cloudinary.CloudinaryInterface, mid midtrans.MidtransService) transaction.TransactionServiceInterface {
	return &TransactionService{
		d:   data,
		cld: cloudinary,
		mt:  mid,
	}
}

func (as *TransactionService) GetTransactions(sort string) ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetAll(sort)
	if err != nil {
		return nil, errors.New("Get All Transactions Failed")
	}
	return result, nil
}

func (as *TransactionService) GetTransaction(id int, sort string) ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetByID(id, sort)
	if err != nil {
		return nil, errors.New("Get By User ID Process Failed")
	}
	return result, nil
}

func (as *TransactionService) GetTransactionByPatientID(id int, sort string) ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetByPatientID(id, sort)
	if err != nil {
		return nil, errors.New("Get By Patient ID Process Failed")
	}
	return result, nil
}

func (as *TransactionService) GetByIDMidtrans(id string) ([]transaction.TransactionInfo, error) {
	result, err := as.d.GetByIDMidtrans(id)
	if err != nil {
		return nil, errors.New("Get By ID Midtrans Process Failed")
	}
	return result, nil
}

func (as *TransactionService) CreateTransaction(newData transaction.Transaction) (*transaction.Transaction, map[string]interface{}, error) {
	totalPrice := newData.PriceMethod + newData.PriceDuration + newData.PriceCounseling

	chargeResp, response, err := as.mt.GenerateTransaction(int(totalPrice), newData.PaymentType)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, nil, errors.New("Generate Transaction Failed")
	}

	newData.MidtransID = chargeResp.OrderID

	result, err := as.d.Insert(newData)

	fmt.Println("Ini new data: ", newData)
	if err != nil {
		return nil, nil, errors.New("Insert Process Failed")
	}
	return result, response, nil
}

func (as *TransactionService) CreateManualTransaction(newData transaction.Transaction) (*transaction.Transaction, error) {

	result, err := as.d.Insert(newData)
	if err != nil {
		return nil, errors.New("Insert Process Failed")
	}
	return result, nil
}

func (as *TransactionService) UpdateTransaction(notificationPayload map[string]interface{}, newData transaction.UpdateTransaction) (bool, error) {
	paymentStatus, orderId, err := as.mt.TransactionStatus(notificationPayload)
	if err != nil {
		return false, errors.New("Transaction Status Failed")
	}

	newData.PaymentStatus = uint(paymentStatus)
	result, err := as.d.GetAndUpdate(newData, orderId)

	if err != nil {
		return false, errors.New("Update Process Failed")
	}

	return result, nil
}

func (as *TransactionService) UpdateTransactionManual(newData transaction.UpdateTransactionManual, id string) (bool, error) {
	if !containsOnlyNumbers(id) {
		result, err := as.d.UpdateWithTrxID(newData, id)

		if err != nil {
			return false, errors.New("Update Process Failed")
		}
		return result, nil

	} else {
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			return false, errors.New("Invalid ID format")
		}
		result, err := as.d.Update(newData, idParsed)

		if err != nil {
			return false, errors.New("Update Process Failed")
		}
		return result, nil

	}
}

func (as *TransactionService) DeleteTransaction(id int) (bool, error) {
	result, err := as.d.Delete(id)

	if err != nil {
		return false, errors.New("Delete Process Failed")
	}

	return result, nil
}

func (as *TransactionService) PaymentProofUpload(newData transaction.PaymentProofDataModel) (string, error) {
	fmt.Println("Ini isi newData: ", newData.PaymentProofPhoto)
	uploadUrl, err := as.cld.UploadImageHelper(newData.PaymentProofPhoto)
	fmt.Println("Ini hasil url: ", uploadUrl)

	if err != nil {
		return "", errors.New("Upload Payment Proof Failed")
	}
	return uploadUrl, nil
}

func containsOnlyNumbers(s string) bool {
	numericRegex := regexp.MustCompile("^[0-9]+$")
	return numericRegex.MatchString(s)
}
