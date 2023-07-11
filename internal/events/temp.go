package events

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/stats/temps"
)

type tempData struct {
	ID   uint    `json:"id"`
	Name string  `json:"name"`
	Temp float32 `json:"temp"`
}

// Temp event structure
type Temp struct {
	Event string
	Data  tempData
}

// NewTempEvent create temp event
func NewTempEvent(model *temps.Model) *Temp {
	data := tempData{
		ID:   model.ID,
		Name: model.Name,
		Temp: model.Data[0].Temp,
	}

	return &Temp{
		Event: FormatEventName(TempTopic, AddTempEvent),
		Data:  data,
	}
}

// GetEvent getting event name
func (e *Temp) GetEvent() string {
	return e.Event
}

// GetData getting data
func (e *Temp) GetData() any {
	return e.Data
}
