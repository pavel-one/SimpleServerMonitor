package tests

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/SimpleServerMonitor/internal/sql"
)

//go:embed mocks/sensors.sql
var InsertSensors string

func GetTestDB() *sqlx.DB {
	db, err := sql.Connect("test")
	if err != nil {
		return nil
	}

	_, err = db.Exec(InsertSensors)
	if err != nil {
		return nil
	}

	return db
}

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
