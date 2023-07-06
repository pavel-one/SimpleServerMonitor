package workers

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/logger"
	"github.com/pavel-one/SimpleServerMonitor/internal/sensors"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
	"time"
)

// SensorWorker worker for writing sensor data
func SensorWorker(period time.Duration, ch events.Chan) error {
	log := logger.NewLogger("SensorWorker")

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
				model, err := rep.AddTemp(sens, chip.Name)
				if err != nil {
					return err
				}

				go func() {
					ch <- events.NewTempEvent(model)
				}()
			}
		}

		time.Sleep(period)
	}
}
