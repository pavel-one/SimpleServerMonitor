package charts

type dataset struct {
	Label    string    `json:"label"`
	SensorID uint      `json:"sensor_id"`
	Data     []float32 `json:"data"`
}

// Chart struct for frontend chart js
type Chart struct {
	Labels   []string   `json:"labels"`
	Datasets []*dataset `json:"datasets"`
}
