package main

import (
	"github.com/pavel-one/sensors/internal/sensors"
	"github.com/pavel-one/sensors/internal/sql"
	"github.com/pavel-one/sensors/internal/ws"
	"time"
)

type App struct {
	ErrorCh chan error
	Ws      *ws.Socket
}

func NewApp() *App {
	w := ws.NewServer(5000, "Socket")
	w.DefaultHandlers()

	return &App{
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
