package packet

import "time"

type Broadcast struct {
	Sender    string
	Type      int
	Content   string
	Timestamp time.Time
}
