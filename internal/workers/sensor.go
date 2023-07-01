package workers

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/Logger"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/sensors"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
	"time"
)

func SensorWorker(period time.Duration, ch events.Chan) error {
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
				model, err := rep.AddTemp(sens)
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
