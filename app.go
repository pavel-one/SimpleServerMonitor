package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/sensors/internal/sensors"
	"time"
)

type App struct {
	DB  *sqlx.DB
	Rep *sensors.SensorRepository
}

func NewApp(DB *sqlx.DB) *App {
	return &App{
		DB:  DB,
		Rep: sensors.NewSensorRepository(DB),
	}
}

// Run Start sensor polling, NOTE: block gorutine runtime
func (a *App) Run(period time.Duration) error {
	for {
		chips, err := sensors.GetChips()
		if err != nil {
			return err
		}

		for _, chip := range chips {
			for _, sens := range chip.Sensors {
				if err := a.Rep.AddTemp(sens); err != nil {
					return err
				}
			}
		}

		time.Sleep(period)
	}
}
