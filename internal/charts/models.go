package charts

type Model struct {
	Time     string  `db:"time"`
	Temp     float32 `db:"temp"`
	Name     string  `db:"name"`
	SensorId uint    `db:"sensor_id"`
}
