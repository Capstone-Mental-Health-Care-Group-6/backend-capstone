package handler

type InputRequest struct {
	Message string `json:"message" validate:"required"`
	Type    string `json:"type"`
}
