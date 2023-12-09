package handler

import "mime/multipart"

type InputRequest struct {
	TopicID           uint           `json:"topic_id" form:"topic_id"`
	PatientID         uint           `json:"patient_id" form:"patient_id"`
	DoctorID          uint           `json:"doctor_id" form:"doctor_id"`
	MethodID          uint           `json:"method_id" form:"method_id"`
	DurationID        uint           `json:"duration_id" form:"duration_id"`
	CounselingID      uint           `json:"counseling_id" form:"counseling_id"`
	UserID            uint           `json:"user_id" form:"user_id"`
	CounselingSession uint           `json:"counseling_session" form:"counseling_session"`
	CounselingType    string         `json:"counseling_type" form:"counseling_type"`
	PriceMethod       uint           `json:"price_method" form:"price_method"`
	PriceDuration     uint           `json:"price_duration" form:"price_duration"`
	PriceCounseling   uint           `json:"price_counseling" form:"price_counseling"`
	PaymentProof      multipart.File `json:"payment_proof" form:"payment_proof"`
	PriceResult       uint           `json:"price_result" form:"price_result"`
	PaymentType       string         `json:"payment_type" form:"payment_type"`
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
