package charts

type dataset struct {
	Label    string    `json:"label"`
	SensorId uint      `json:"sensor_id"`
	Data     []float32 `json:"data"`
}

type Chart struct {
	Labels   []string   `json:"labels"`
	Datasets []*dataset `json:"datasets"`
}
