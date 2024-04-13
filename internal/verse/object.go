package verse

import "time"

type Verse struct {
	Created time.Time `json:"created,omitempty"`
	Address string    `json:"address,omitempty"`
	Content string    `json:"content,omitempty"`
	Id      int       `json:"id,omitempty"`
}
