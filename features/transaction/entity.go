package transaction

import "github.com/labstack/echo/v4"

type Transaction struct {
	TopicID      uint   `json:"topic_id"`
	PatientID    uint   `json:"patient_id"`
	DoctorID     uint   `json:"doctor_id"`
	MethodID     uint   `json:"method_id"`
	DurationID   uint   `json:"duration_id"`
	CounselingID uint   `json:"counseling_id"`
	UserID       uint   `json:"user_id"`
	MidtransID   string `json:"midtrans_id"`

	CounselingSession uint   `json:"counseling_session"`
	CounselingType    string `json:"counseling_type"`

	PriceMethod     uint `json:"price_method"`
	PriceDuration   uint `json:"price_duration"`
	PriceCounseling uint `json:"price_counseling"`
	PriceResult     uint `json:"price_result"`

	PaymentStatus uint   `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
}

type TransactionInfo struct {
	ID             uint   `json:"id"`
	TopicName      uint   `json:"topic_name"`
	PatientName    uint   `json:"patient_name"`
	DoctorName     uint   `json:"doctor_name"`
	MethodName     uint   `json:"method_name"`
	DurationName   uint   `json:"duration_name"`
	CounselingName uint   `json:"counseling_name"`
	UserName       uint   `json:"user_name"`
	MidtransID     string `json:"midtrans_id"`

	CounselingSession uint   `json:"counseling_session"`
	CounselingType    string `json:"counseling_type"`

	PriceMethod     uint `json:"price_method"`
	PriceDuration   uint `json:"price_duration"`
	PriceCounseling uint `json:"price_counseling"`
	PriceResult     uint `json:"price_result"`

	PaymentStatus uint   `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
}

type UpdateTransaction struct {
	PaymentStatus uint `json:"payment_status"`
}

type TransactionHandlerInterface interface {
	GetTransactions() echo.HandlerFunc
	GetTransaction() echo.HandlerFunc
	CreateTransaction() echo.HandlerFunc
	NotifTransaction() echo.HandlerFunc
}

type TransactionServiceInterface interface {
	GetTransactions() ([]TransactionInfo, error)
	GetTransaction(id int) ([]TransactionInfo, error)
	CreateTransaction(newData Transaction) (*Transaction, error)
	UpdateTransaction(newData UpdateTransaction, id string) (bool, error)
}

type TransactionDataInterface interface {
	GetAll() ([]TransactionInfo, error)
	GetByID(id int) ([]TransactionInfo, error)
	Insert(newData Transaction) (*Transaction, error)
	// Update(newData Transaction, id int) (bool, error)
	GetAndUpdate(newData UpdateTransaction, id string) (bool, error)
}