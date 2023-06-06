package sensors

import "errors"

type Sensor struct {
	Name     string  `json:"name"`
	Temp     float32 `json:"temp"`
	HighTemp float32 `json:"high_temp"`
	CritTemp float32 `json:"crit_temp"`
	RawData  string  `json:"raw_data"`
}

func NewSensor(name string, rawData string) (*Sensor, error) {
	if !StrHasTemp(rawData) {
		return nil, errors.New("sensor raw data not has temp")
	}

	sensor := &Sensor{
		Name:    name,
		RawData: rawData,
	}

	temps := StrExtractTemps(rawData)

	switch len(temps) {
	case 1:
		sensor.Temp = temps[0]
		break
	case 2:
		sensor.Temp = temps[0]
		sensor.CritTemp = temps[1]
		break
	case 3:
		sensor.Temp = temps[0]
		sensor.HighTemp = temps[1]
		sensor.CritTemp = temps[2]
	}

	return sensor, nil
}

type Chip struct {
	Name    string    `json:"name"`
	Sensors []*Sensor `json:"sensors"`
}
