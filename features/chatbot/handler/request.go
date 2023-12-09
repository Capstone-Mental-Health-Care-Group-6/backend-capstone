package handler

type InputRequest struct {
	Prompt string `json:"prompt" form:"prompt" validate:"required"`
}
