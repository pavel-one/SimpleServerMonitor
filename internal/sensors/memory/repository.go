package memory

import (
	"fmt"
	"github.com/pavel-one/SimpleServerMonitor/internal/base"
	"github.com/pavel-one/SimpleServerMonitor/internal/db"
	"time"
)

var schema = `
CREATE TABLE IF NOT EXISTS Memory
(
    key     VARCHAR NOT NULL,
    temp 	INTEGER NOT NULL,
    time 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type Model struct {
	Key  string    `db:"key"`
	Temp float64   `db:"temp"`
	Time time.Time `db:"time"`
}

type Repository base.Repository

func NewRepository() (*Repository, error) {
	connection, err := db.DefaultConnection()
	if err != nil {
		return nil, err
	}

	if _, err := connection.Exec(schema); err != nil {
		return nil, fmt.Errorf("error create schema: %s", err)
	}

	return &Repository{DB: connection}, nil
}
