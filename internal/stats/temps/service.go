package temps

import (
	"github.com/ssimunic/gosensors"
)

// GetChips getting all chips
func GetChips() ([]Chip, error) {
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
		return nil, err
	}

	return out, nil
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
