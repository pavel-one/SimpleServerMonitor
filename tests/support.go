package tests

import (
	_ "embed" // embed sql file
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
)

//go:embed mocks/sensors.sql
var sensorsData string // sensors data

// GetTestDB getting testing database with data
func GetTestDB() *sqlx.DB {
	db, err := sql.Connect("test")
	if err != nil {
		return nil
	}

	_, err = db.Exec(sensorsData)
	if err != nil {
		return nil
	}

	return db
}

// GetEmptyTestDB getting empty testing database
func GetEmptyTestDB() *sqlx.DB {
	db, err := sql.Connect("test")
	if err != nil {
		return nil
	}

	_, err = db.Exec("DELETE FROM chips;\nVACUUM;\nDELETE FROM sensors;\nVACUUM;\nDELETE FROM sensors_data;\nVACUUM;")
	if err != nil {
		return nil
	}

	return db
}
