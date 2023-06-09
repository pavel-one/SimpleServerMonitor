package main

import (
	"github.com/pavel-one/sensors/internal/sensors"
	"github.com/pavel-one/sensors/internal/sql"
	"log"
)

func main() {
	db, err := sql.Connect("test")
	if err != nil {
		log.Fatalln(err)
	}

	rep := sensors.NewSensorRepository(db)
	find, err := rep.Find(1)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(find)
}
