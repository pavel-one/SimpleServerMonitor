package main

import (
	"github.com/pavel-one/sensors/internal/events"
	"github.com/pavel-one/sensors/internal/workers"
	"github.com/pavel-one/sensors/internal/ws"
	"time"
)

type App struct {
	ErrorCh chan error
	EventCh events.Chan
	Ws      *ws.Socket
}

func NewApp() *App {
	ech := make(events.Chan)

	w := ws.NewServer(5000, "Socket", ech)
	w.SetDefault()

	return &App{
		ErrorCh: make(chan error, 1),
		EventCh: ech,
		Ws:      w,
	}
}

// Run all application components
func (a *App) Run() error {
	go func(ch chan<- error) {
		if err := workers.SensorWorker(time.Second * 5); err != nil {
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
