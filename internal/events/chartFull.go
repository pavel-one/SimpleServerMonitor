package events

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/charts"
)

type ChartFull struct {
	Event string        `json:"event"`
	Data  *charts.Chart `json:"data"`
}

func NewChartFull(data *charts.Chart) *ChartFull {
	return &ChartFull{
		Event: FormatEventName(TempTopic, LoadTempEvent),
		Data:  data,
	}
}

func (e *ChartFull) GetEvent() string {
	return e.Event
}

func (e *ChartFull) GetData() any {
	return e.Data
}
