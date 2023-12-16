package data

import "gorm.io/gorm"

type Transaction struct {
	*gorm.Model
	TopicID      uint   `gorm:"default:null;column:topic_id"`
	PatientID    uint   `gorm:"default:null;column:patient_id"`
	DoctorID     uint   `gorm:"default:null;column:doctor_id"`
	MethodID     uint   `gorm:"default:null;column:method_id"`
	DurationID   uint   `gorm:"default:null;column:duration_id"`
	CounselingID uint   `gorm:"default:null;column:counseling_id"`
	UserID       uint   `gorm:"default:null;column:user_id"`
	MidtransID   string `gorm:"column:midtrans_id"`

	CounselingSession uint   `gorm:"default:null;column:counseling_session"`
	CounselingType    string `gorm:"default:null;column:counseling_type;type:enum('A','B','C')"`

	PriceMethod     uint `gorm:"default:null;column:price_method"`
	PriceDuration   uint `gorm:"default:null;column:price_duration"`
	PriceCounseling uint `gorm:"default:null;column:price_counseling"`
	PriceResult     uint `gorm:"column:price_result"`

	PaymentProof  string `gorm:"default:null;column:payment_proof;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"`
	PaymentStatus uint   `gorm:"column:payment_status"`
	PaymentType   string `gorm:"column:payment_type;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"`
}

func (Transaction) TableName() string {
	return "transactions"
}
