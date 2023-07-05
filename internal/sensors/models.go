package sensors

import "time"

type DataItem struct {
	Temp      float32   `json:"temp" db:"temp"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Model struct {
	ID       uint       `json:"id" db:"id"`
	Name     string     `json:"name" db:"name"`
	HighTemp float32    `json:"high_temp" db:"high_temp"`
	CritTemp float32    `json:"crit_temp" db:"crit_temp"`
	ChipId   uint       `json:"chip_id" db:"chip_id"`
	Data     []DataItem `json:"data" db:"-"`
}

type ChipModel struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
}
