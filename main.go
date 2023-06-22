package main

import (
	"github.com/pavel-one/sensors/internal/Logger"
)

var log = Logger.NewLogger("Application")

func main() {
	app := NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("Application fatal error: %s", err)
	}
}
