package charts

// Model database model for chart
type Model struct {
	Time     string  `db:"time"`
	Temp     float32 `db:"temp"`
	Name     string  `db:"name"`
	SensorID uint    `db:"sensor_id"`
}
