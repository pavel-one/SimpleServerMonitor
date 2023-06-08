package sensors

import "github.com/jmoiron/sqlx"

type Model struct {
	ID       uint      `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	HighTemp float32   `json:"high_temp" db:"high_temp"`
	CritTemp float32   `json:"crit_temp" db:"crit_temp"`
	Data     []float32 `json:"data"`
}

type SensorRepository struct {
	DB *sqlx.DB
}
