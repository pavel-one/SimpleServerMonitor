package main

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/workers"
	"github.com/pavel-one/SimpleServerMonitor/internal/ws"
	"time"
)

type App struct {
	ErrorCh chan error
	EventCh events.Chan
	Ws      *ws.Socket
}

func NewApp() *App {
	ech := make(events.Chan, 1)

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
		if err := workers.SensorWorker(time.Second*5, a.EventCh); err != nil {
			log.Errorln("Sensor worker is failed", err)
			ch <- err
		}
	}(a.ErrorCh)

	go func(ch chan<- error) {
		if err := a.Ws.Run(); err != nil {
			log.Errorln("Websocket server is failed", err)
			ch <- err
		}
	}(a.ErrorCh)

	go func(ch chan<- error) {
		if err := workers.WebsocketWorker(a.Ws.Server, a.EventCh); err != nil {
			log.Errorln("Sensor worker is failed", err)
			ch <- err
		}
	}(a.ErrorCh)

	return <-a.ErrorCh
}
