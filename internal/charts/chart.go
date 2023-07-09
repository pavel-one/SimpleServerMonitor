package charts

import "time"

type dataset struct {
	Name     string  `json:"name"`
	SensorID uint    `json:"sensor_id"`
	Data     [][]any `json:"data"`
}

// Chart struct for frontend chart js
type Chart struct {
	DateStart time.Time  `json:"date_start"`
	Datasets  []*dataset `json:"datasets"`
}
