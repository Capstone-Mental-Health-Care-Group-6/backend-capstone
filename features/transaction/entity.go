package transaction

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	TopicID      uint   `json:"topic_id"`
	PatientID    uint   `json:"patient_id"`
	DoctorID     uint   `json:"doctor_id"`
	MethodID     uint   `json:"method_id"`
	DurationID   uint   `json:"duration_id"`
	CounselingID uint   `json:"counseling_id"`
	UserID       uint   `json:"user_id"`
	MidtransID   string `json:"transaction_id"`

	CounselingSession uint   `json:"counseling_session"`
	CounselingType    string `json:"counseling_type"`

	PriceMethod     uint `json:"price_method"`
	PriceDuration   uint `json:"price_duration"`
	PriceCounseling uint `json:"price_counseling"`
	PriceResult     uint `json:"price_result"`

	PaymentProof  string `json:"payment_proof"`
	PaymentStatus uint   `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
}

type TransactionInfo struct {
	UserID        uint   `json:"user_id"`
	MidtransID    string `json:"transaction_id"`
	PriceResult   uint   `json:"price_result"`
	PaymentStatus uint   `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
}

type PaymentProofDataModel struct {
	PaymentProofPhoto multipart.File `json:"payment_proof"`
}

type UpdateTransactionManual struct {
	UserID          uint   `json:"user_id"`
	PriceMethod     uint   `json:"price_method"`
	PriceDuration   uint   `json:"price_duration"`
	PriceCounseling uint   `json:"price_counseling"`
	PriceResult     uint   `json:"price_result"`
	PaymentStatus   uint   `json:"payment_status"`
	PaymentType     string `json:"payment_type"`
}

type UpdateTransaction struct {
	PaymentStatus uint `json:"payment_status"`
}

type TransactionHandlerInterface interface {
	GetTransactions() echo.HandlerFunc
	GetTransaction() echo.HandlerFunc
	GetTransactionByMidtransID() echo.HandlerFunc
	CreateTransaction() echo.HandlerFunc
	NotifTransaction() echo.HandlerFunc
	DeleteTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
}

type TransactionServiceInterface interface {
	GetTransactions(sort string) ([]TransactionInfo, error)
	// GetTransactionsSort(sort string) ([]TransactionInfo, error)
	GetTransaction(id int, sort string) ([]Transaction, error)
	CreateTransaction(newData Transaction) (*Transaction, map[string]interface{}, error)
	CreateManualTransaction(newData Transaction) (*Transaction, error)
	GetByIDMidtrans(id string) ([]TransactionInfo, error)
	UpdateTransaction(notificationPayload map[string]interface{}, newData UpdateTransaction) (bool, error)
	DeleteTransaction(id int) (bool, error)
	UpdateTransactionManual(newData UpdateTransactionManual, id string) (bool, error)
	PaymentProofUpload(newData PaymentProofDataModel) (string, error)
}

type TransactionDataInterface interface {
	GetAll(sort string) ([]TransactionInfo, error)
	// GetAllSort(sort string) ([]TransactionInfo, error)
	GetByID(id int, sort string) ([]Transaction, error)
	GetByIDMidtrans(id string) ([]TransactionInfo, error)
	Insert(newData Transaction) (*Transaction, error)
	Update(newData UpdateTransactionManual, id int) (bool, error)
	UpdateWithTrxID(newData UpdateTransactionManual, id string) (bool, error)
	GetAndUpdate(newData UpdateTransaction, id string) (bool, error)
	Delete(id int) (bool, error)
}
