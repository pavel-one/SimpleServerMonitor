package temps

import "time"

func Worker(t time.Duration) error {
	rep, err := NewStatRepository()
	if err != nil {
		return err
	}

	for {
		sensors, err := GetSensors()
		if err != nil {
			return err
		}

		stats, err := GetStats(sensors)
		if err != nil {
			return err
		}

		for _, s := range stats {
			if err := rep.Save(s); err != nil {
				return err
			}
		}

		time.Sleep(t)
	}
}
