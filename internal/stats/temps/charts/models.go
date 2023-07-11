package charts

import "time"

// Model database model for chart
type Model struct {
	Time     time.Time `db:"time"`
	Temp     float32   `db:"temp"`
	Name     string    `db:"name"`
	SensorID uint      `db:"sensor_id"`
}
