package sql

import (
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
)

//go:embed memory.sql
var schema string

// Connect to database
func Connect(dbname string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", fmt.Sprintf("./%s.sqlite3", dbname))
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
