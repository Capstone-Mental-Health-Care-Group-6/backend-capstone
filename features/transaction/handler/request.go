package handler

type InputRequest struct {
	TopicID           uint   `json:"topic_id"`
	PatientID         uint   `json:"patient_id"`
	DoctorID          uint   `json:"doctor_id"`
	MethodID          uint   `json:"method_id"`
	DurationID        uint   `json:"duration_id"`
	CounselingID      uint   `json:"counseling_id"`
	UserID            uint   `json:"user_id"`
	CounselingSession uint   `json:"counseling_session"`
	CounselingType    string `json:"counseling_type"`
	PriceMethod       uint   `json:"price_method"`
	PriceDuration     uint   `json:"price_duration"`
	PriceCounseling   uint   `json:"price_counseling"`
	PriceResult       uint   `json:"price_result"`
	PaymentType       string `json:"payment_type"`
}
