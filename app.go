package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/pavel-one/SimpleServerMonitor/internal/events"
	"github.com/pavel-one/SimpleServerMonitor/internal/workers"
	"github.com/pavel-one/SimpleServerMonitor/internal/ws"
	"io/fs"
	"net/http"
	"os"
	"time"
)

//go:embed frontend/dist/*
var frontendFS embed.FS

// App application wrapper for run worker and services
type App struct {
	ErrorCh chan error
	EventCh events.Chan
	Ws      *ws.Socket
	Router  *gin.Engine
}

// NewApp create new application
func NewApp() *App {
	ech := make(events.Chan, 1)

	w := ws.NewServer(5000, "Socket", ech)
	w.SetDefault()

	router := gin.Default()

	sub, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		panic(err)
	}

	router.StaticFS("/", http.FS(sub))

	return &App{
		ErrorCh: make(chan error, 1),
		EventCh: ech,
		Ws:      w,
		Router:  router,
	}
}

// Run all application components
func (a *App) Run() error {

	// sensor worker
	go func(ch chan<- error) {
		if err := workers.SensorWorker(time.Second*5, a.EventCh); err != nil {
			log.Errorln("Sensor worker is failed", err)
			ch <- err
		}
	}(a.ErrorCh)

	// ws server
	go func(ch chan<- error) {
		if err := a.Ws.Run(); err != nil {
			log.Errorln("Websocket server is failed", err)
			ch <- err
		}
	}(a.ErrorCh)

	// ws worker
	go func(ch chan<- error) {
		if err := workers.WebsocketWorker(a.Ws.Server, a.EventCh); err != nil {
			log.Errorln("Sensor worker is failed", err)
			ch <- err
		}
	}(a.ErrorCh)

	//http server
	go func(ch chan<- error) {
		r := os.Getenv("HTTP_SERVER")
		if r == "false" {
			return
		}
		log.Infoln("Starting http server on port: 8080")
		if err := a.Router.Run("127.0.0.1:8080"); err != nil {
			log.Errorln("Http server is failed:", err)
			ch <- err
		}
	}(a.ErrorCh)

	return <-a.ErrorCh
}
