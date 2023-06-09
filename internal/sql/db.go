package sql

import (
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed db/schema.sql
var schema string

func Connect(dbname string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", fmt.Sprintf("./%s.sqlite3", dbname))
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}