package events

// AddTempEvent event for adding temperature
// LoadTempEvent event for load all temperature chips and sensors
const (
	AddTempEvent  = "add"
	LoadTempEvent = "all"
)

// TempTopic for temperature data
const (
	TempTopic = "temp" // topic for temperature data
)
