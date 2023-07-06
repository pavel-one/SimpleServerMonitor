package main

import (
	"github.com/pavel-one/SimpleServerMonitor/internal/logger"
)

var log = logger.NewLogger("Application")

func main() {
	app := NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("Application fatal error: %s", err)
	}
}
