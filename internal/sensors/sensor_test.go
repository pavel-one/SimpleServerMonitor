package sensors

import "testing"

func TestNewSensor(t *testing.T) {
	name := "Core 3"
	rawData := "+46.0°C  (high = +80.0°C, crit = +100.0°C)"

	sensor, err := NewSensor(name, rawData)

	if err != nil {
		t.Fatal(err)
	}

	if sensor.Temp != 46.0 {
		t.Fatalf("sensor not correct temp: %f", sensor.Temp)
	}

	if sensor.CritTemp != 100.0 {
		t.Fatalf("sensor not correct crit temp: %f", sensor.CritTemp)
	}

	if sensor.HighTemp != 80.0 {
		t.Fatalf("sensor not correct high temp: %f", sensor.HighTemp)
	}
}

func TestNewSensorError(t *testing.T) {
	name := "Core 3"
	rawData := "Is not temp"

	_, err := NewSensor(name, rawData)

	if err == nil {
		t.Fatal("not getting error for fake sensor")
	}
}

func TestNewSensorWithoutMaxTemps(t *testing.T) {
	name := "Core 3"
	rawData := "+46.0°C"

	sensor, err := NewSensor(name, rawData)

	if err != nil {
		t.Fatal(err)
	}

	if sensor.Temp == 0 {
		t.Fatal("not getting temp")
	}

	if sensor.HighTemp != 0 || sensor.CritTemp != 0 {
		t.Fatal("not correct getting temp")
	}
}
