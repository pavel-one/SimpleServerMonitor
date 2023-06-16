package main

import (
	"github.com/pavel-one/sensors/internal/Logger"
	"github.com/pavel-one/sensors/internal/sql"
)

var log = Logger.NewLogger("Application")

func main() {
	db, err := sql.Connect("db")
	if err != nil {
		log.Fatalln(err)
	}

	app := NewApp(db)

	if err := app.Run(); err != nil {
		log.Fatalf("Application fatal error: %s", err)
	}
}
