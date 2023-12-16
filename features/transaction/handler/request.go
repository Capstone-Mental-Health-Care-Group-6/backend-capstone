package handler

type InputRequest struct {
	TopicID           uint   `json:"topic_id" form:"topic_id" validate:"required"`
	PatientID         uint   `json:"patient_id" form:"patient_id" validate:"required"`
	DoctorID          uint   `json:"doctor_id" form:"doctor_id" validate:"required"`
	MethodID          uint   `json:"method_id" form:"method_id" validate:"required"`
	DurationID        uint   `json:"duration_id" form:"duration_id" validate:"required"`
	CounselingID      uint   `json:"counseling_id" form:"counseling_id" validate:"required"`
	UserID            uint   `json:"user_id" form:"user_id" validate:"required"`
	CounselingSession uint   `json:"counseling_session" form:"counseling_session" validate:"required"`
	CounselingType    string `json:"counseling_type" form:"counseling_type" validate:"required"`
	PriceMethod       uint   `json:"price_method" form:"price_method" validate:"required"`
	PriceDuration     uint   `json:"price_duration" form:"price_duration" validate:"required"`
	PriceCounseling   uint   `json:"price_counseling" form:"price_counseling" validate:"required"`
	PriceResult       uint   `json:"price_result" form:"price_result" validate:"required"`
	PaymentType       string `json:"payment_type" form:"payment_type" validate:"required"`
}

type UpdateRequest struct {
	UserID          uint   `json:"user_id" form:"user_id"`
	PriceMethod     uint   `json:"price_method" form:"price_method"`
	PriceDuration   uint   `json:"price_duration" form:"price_duration"`
	PriceCounseling uint   `json:"price_counseling" form:"price_counseling"`
	PriceResult     uint   `json:"price_result" form:"price_result"`
	PaymentStatus   uint   `json:"payment_status" form:"payment_status"`
	PaymentType     string `json:"payment_type" form:"payment_type"`
}
