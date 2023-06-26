package workers

import (
	"github.com/pavel-one/sensors/internal/Logger"
	"github.com/pavel-one/sensors/internal/sensors"
	"github.com/pavel-one/sensors/internal/sql"
	"time"
)

func SensorWorker(period time.Duration) error {
	log := Logger.NewLogger("SensorWorker")

	db, err := sql.Connect("db")
	if err != nil {
		return err
	}
	rep := sensors.NewSensorRepository(db)

	log.Infoln("Run temps pulling")

	for {
		chips, err := sensors.GetChips()
		if err != nil {
			return err
		}

		for _, chip := range chips {
			for _, sens := range chip.Sensors {
				if err := rep.AddTemp(sens); err != nil {
					return err
				}
			}
		}

		time.Sleep(period)
	}
}