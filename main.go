package main

import (
	"github.com/pavel-one/sensors/internal/sql"
	"log"
	"time"
)

func main() {
	db, err := sql.Connect("db")
	if err != nil {
		log.Fatalln(err)
	}

	app := NewApp(db)

	if err := app.Run(time.Second * 5); err != nil {
		log.Fatalln(err)
	}
}
