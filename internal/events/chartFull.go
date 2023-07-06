package events

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/charts"
)

// ChartFull event load all temperature
type ChartFull struct {
	Event string        `json:"event"`
	Data  *charts.Chart `json:"data"`
}

// NewChartFull create event for load all temperature
func NewChartFull(data *charts.Chart) *ChartFull {
	return &ChartFull{
		Event: FormatEventName(TempTopic, LoadTempEvent),
		Data:  data,
	}
}

// GetEvent get event
func (e *ChartFull) GetEvent() string {
	return e.Event
}

// GetData get data
func (e *ChartFull) GetData() any {
	return e.Data
}
