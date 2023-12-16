package dto

type Request struct {
	Text string `json:"text"`
	Blob string `json:"blob,omitempty"`
	// Sender int    `json:"sender"`
}
