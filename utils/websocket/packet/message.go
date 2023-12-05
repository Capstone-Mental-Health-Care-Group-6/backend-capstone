package packet

import "time"

type Message struct {
	Room int       `json:"room"`
	From int       `json:"from"`
	To   int       `json:"to"`
	Text string    `json:"text"`
	Blob string    `json:"blob"`
	Time time.Time `json:"time"`
}
