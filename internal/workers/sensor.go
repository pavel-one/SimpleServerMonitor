package workers

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/logger"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
	"github.com/pavel-one/SimpleServerMonitor/internal/stats/temps"
	"github.com/pavel-one/SimpleServerMonitor/internal/stats/temps/charts"
	"time"
)

// SensorWorker worker for writing sensor data
func SensorWorker(period time.Duration, ch events.Chan) error {
	log := logger.NewLogger("SensorWorker")

	db, err := sql.Connect("db")
	if err != nil {
		return err
	}
	rep := temps.NewSensorRepository(db)
	chartRep := charts.NewRepository(db)

	log.Infoln("Run temps pulling")

	for {
		chips, err := temps.GetChips()
		if err != nil {
			return err
		}

		for _, chip := range chips {
			for _, sens := range chip.Sensors {
				if _, err := rep.AddTemp(sens, chip.Name); err != nil {
					return err
				}
			}
		}

		go func() {
			last, err := chartRep.GetLast(charts.TypeSecond)
			if err != nil {
				log.Errorln("Error load last charts:", err)
				return
			}

			ch <- events.NewChart(last, events.AddTempEvent)
		}()
		time.Sleep(period)
	}
}
