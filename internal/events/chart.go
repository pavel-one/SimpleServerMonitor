package events

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/stats/temps/charts"
)

// Chart event load all temperature
type Chart struct {
	Event string        `json:"event"`
	Data  *charts.Chart `json:"data"`
}

// NewChart create event for load all temperature
func NewChart(data *charts.Chart, eventName string) *Chart {
	return &Chart{
		Event: FormatEventName(TempTopic, eventName),
		Data:  data,
	}
}

// GetEvent get event
func (e *Chart) GetEvent() string {
	return e.Event
}

// GetData get data
func (e *Chart) GetData() any {
	return e.Data
}
