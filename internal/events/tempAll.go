package events

import "github.com/pavel-one/SimpleServerMonitor/internal/sensors"

type TempAll struct {
	Event string           `json:"event"`
	Data  []*sensors.Model `json:"data"`
}

func NewTempAll(data []*sensors.Model) *TempAll {
	return &TempAll{
		Event: FormatEventName(TempTopic, LoadTempEvent),
		Data:  data,
	}
}

func (e *TempAll) GetEvent() string {
	return e.Event
}

func (e *TempAll) GetData() any {
	return e.Data
}
