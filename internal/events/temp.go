package events

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/sensors"
)

type tempData struct {
	ID   uint    `json:"id"`
	Name string  `json:"name"`
	Temp float32 `json:"temp"`
}

type Temp struct {
	Event string
	Data  tempData
}

func NewTempEvent(model *sensors.Model) *Temp {
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

func (e *Temp) GetEvent() string {
	return e.Event
}

func (e *Temp) GetData() any {
	return e.Data
}
