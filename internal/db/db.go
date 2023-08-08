package db

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func DefaultConnection() (*sqlx.DB, error) {
	return Connect("db")
}

func Connect(dbname string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("sqlite", dbname+".sqlite3")
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
