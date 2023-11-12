package transaction

import "github.com/labstack/echo/v4"

type Transaction struct {
	TopicID      *uint   `json:"topic_id"`
	PatientID    *uint   `json:"patient_id"`
	DoctorID     *uint   `json:"doctor_id"`
	MethodID     *uint   `json:"method_id"`
	DurationID   *uint   `json:"duration_id"`
	CounselingID *uint   `json:"counseling_id"`
	UserID       int     `json:"user_id"`
	MidtransID   *string `json:"midtrans_id"`

	CounselingSession *int    `json:"counseling_session"`
	CounselingType    *string `json:"counseling_type"`

	PriceMethod     *int `json:"price_method"`
	PriceDuration   *int `json:"price_duration"`
	PriceCounseling *int `json:"price_counseling"`
	PriceResult     *int `json:"price_result"`

	PaymentStatus int    `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
}

type TransactionInfo struct {
	ID             uint   `json:"id"`
	TopicName      *uint  `json:"topic_name"`
	PatientName    *uint  `json:"patient_name"`
	DoctorName     *uint  `json:"doctor_name"`
	MethodName     *uint  `json:"method_name"`
	DurationName   *uint  `json:"duration_name"`
	CounselingName *uint  `json:"counseling_name"`
	UserName       uint   `json:"user_name"`
	MidtransID     string `json:"midtrans_id"`

	CounselingSession *int   `gorm:"column:counseling_session" json:"counseling_session"`
	CounselingType    string `gorm:"column:counseling_type;type:enum('A','B','C')" json:"counseling_type"`

	PriceMethod     int `gorm:"column:price_method" json:"price_method"`
	PriceDuration   int `gorm:"column:price_duration" json:"price_duration"`
	PriceCounseling int `gorm:"column:price_counseling" json:"price_counseling"`
	PriceResult     int `gorm:"column:price_result" json:"price_result"`

	PaymentStatus int    `gorm:"column:payment_status" json:"payment_status"`
	PaymentType   string `gorm:"column:payment_type" json:"payment_type"`
}

type TransactionHandlerInterface interface {
	GetTransactions() echo.HandlerFunc
	GetTransaction() echo.HandlerFunc
	CreateTransaction() echo.HandlerFunc
}

type TransactionServiceInterface interface {
	GetTransactions() ([]TransactionInfo, error)
	GetTransaction(id int) ([]TransactionInfo, error)
	CreateTransaction(newData Transaction) (*Transaction, error)
}

type TransactionDataInterface interface {
	GetAll() ([]TransactionInfo, error)
	GetByID(id int) ([]TransactionInfo, error)
	Insert(newData Transaction) (*Transaction, error)
}
