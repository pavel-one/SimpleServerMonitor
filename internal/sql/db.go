package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB sqlx.DB

var schema = `
CREATE TABLE IF NOT EXISTS sensors (
    id integer
        constraint sensors_pk
            primary key,
    name varchar not null
);

CREATE TABLE IF NOT EXISTS sensors_data (
    temp DECIMAL(2,2) not null,
    sensor_id integer not null,
    FOREIGN KEY(sensor_id) REFERENCES sensors(id)
);
`

func GetDatabase() (DB, error) {
	db, err := sqlx.Connect("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	exec := db.MustExec(schema)

	if db {

	}
}

func (d *DB) Init() error {

}
