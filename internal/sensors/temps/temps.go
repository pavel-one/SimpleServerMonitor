package temps

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/logger"
	"github.com/shirou/gopsutil/v3/host"
)

var log = logger.NewLogger("temp")

type Stat struct {
	Temp float64
	Key  string
}

func GetStats(temperatures []host.TemperatureStat) ([]*Stat, error) {
	out := make([]*Stat, 0, len(temperatures))

	for _, v := range temperatures {

		find, index := duplicateSensor(v.SensorKey, out)

		if find {
			out[index].Temp = (out[index].Temp + v.Temperature) / 2
			continue
		}

		out = append(out, &Stat{
			Temp: v.Temperature,
			Key:  v.SensorKey,
		})

	}

	return out, nil
}

func GetSensors() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}
