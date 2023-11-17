package packet

import "time"

type Message struct {
	From      int
	To        int
	Text      string
	Image     []byte
	Timestamp time.Time
}
