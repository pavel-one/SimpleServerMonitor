package memory

import "time"

type Model struct {
	Percent   float64   `json:"percent" db:"percent"`
	Free      uint64    `json:"free" db:"free"`
	Total     uint64    `json:"total" db:"total"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
