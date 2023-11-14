package data

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model
	TopicID      *uint  `gorm:"column:topic_id"`
	PatientID    *uint  `gorm:"column:patient_id"`
	DoctorID     *uint  `gorm:"column:doctor_id"`
	MethodID     *uint  `gorm:"column:method_id"`
	DurationID   *uint  `gorm:"column:duration_id"`
	CounselingID *uint  `gorm:"column:counseling_id"`
	UserID       int    `gorm:"column:user_id"`
	MidtransID   string `gorm:"column:midtrans_id"`

	CounselingSession *int    `gorm:"column:counseling_session"`
	CounselingType    *string `gorm:"column:counseling_type;type:enum('A','B','C')"`

	PriceMethod     *int `gorm:"column:price_method"`
	PriceDuration   *int `gorm:"column:price_duration"`
	PriceCounseling *int `gorm:"column:price_counseling"`
	PriceResult     int  `gorm:"column:price_result"`

	PaymentStatus int    `gorm:"column:payment_status"`
	PaymentType   string `gorm:"column:payment_type"`
}
