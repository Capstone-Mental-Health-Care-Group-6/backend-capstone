package dto

type Request struct {
	Sender int    `json:"sender"`
	Text   string `json:"text"`
	Blob   string `json:"blob,omitempty"`
}
