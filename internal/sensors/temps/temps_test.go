package temps

import (
	"github.com/shirou/gopsutil/v3/host"
	"testing"
)

func TestGetStats(t *testing.T) {
	temperatures := []host.TemperatureStat{
		host.TemperatureStat{
			SensorKey:   "Test1",
			Temperature: 25.0,
			High:        25,
			Critical:    65,
		},
		host.TemperatureStat{
			SensorKey:   "Test2",
			Temperature: 22.0,
			High:        25,
			Critical:    65,
		},
		host.TemperatureStat{
			SensorKey:   "Test1",
			Temperature: 25.0,
			High:        25,
			Critical:    65,
		},
		host.TemperatureStat{
			SensorKey:   "Test1",
			Temperature: 23.0,
			High:        25,
			Critical:    65,
		},
	}

	stats, err := GetStats(temperatures)
	if err != nil {
		t.Fatal(err)
	}

	if stats[0].Temp != 24.5 {
		t.Fatalf("Temp %s != 25: %f", stats[0].Key, stats[0].Temp)
	}

	if stats[1].Temp != 22 {
		t.Fatalf("Temp %s != 22: %f", stats[1].Key, stats[1].Temp)
	}
}
