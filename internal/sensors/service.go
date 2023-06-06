package sensors

import (
	"fmt"
	"github.com/ssimunic/gosensors"
)

func GetSensors() []Chip {
	sensors, err := gosensors.NewFromSystem()
	out := make([]Chip, 0)

	for name, chip := range sensors.Chips {
		if len(chip) == 0 {
			continue
		}

		sensors := getSensors(chip)

		out = append(out, Chip{
			Name:    name,
			Sensors: sensors,
		})
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(sensors)

	return out
}

func getSensors(ent gosensors.Entries) []*Sensor {
	sensors := make([]*Sensor, 0)

	for sName, sData := range ent {
		s, err := NewSensor(sName, sData)
		if err != nil {
			continue
		}

		sensors = append(sensors, s)
	}

	return sensors
}
