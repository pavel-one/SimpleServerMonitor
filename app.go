package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/sensors/internal/sensors"
	"github.com/pavel-one/sensors/internal/ws"
	"time"
)

type App struct {
	DB      *sqlx.DB
	Rep     *sensors.SensorRepository
	ErrorCh chan error
	Ws      *ws.Socket
}

func NewApp(DB *sqlx.DB) *App {
	w := ws.NewServer(5000, "Socket")
	w.DefaultHandlers()

	return &App{
		DB:      DB,
		Rep:     sensors.NewSensorRepository(DB),
		ErrorCh: make(chan error, 1),
		Ws:      w,
	}
}

// Run all application components
func (a *App) Run() error {
	go func(ch chan<- error) {
		if err := a.RunTemps(time.Second * 5); err != nil {
			ch <- err
		}
	}(a.ErrorCh)

	go func(ch chan<- error) {
		if err := a.Ws.Run(); err != nil {
			ch <- err
		}
	}(a.ErrorCh)

	return <-a.ErrorCh
}

// RunTemps Start sensor polling
func (a *App) RunTemps(period time.Duration) error {
	log.Infoln("Run temps pulling")

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
