package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS sensors (
    id integer
        constraint sensors_pk
            primary key,
    name varchar not null,
    high_temp DECIMAL(2,2) not null,
    crit_temp DECIMAL(2,2) not null
);

CREATE TABLE IF NOT EXISTS sensors_data (
    temp DECIMAL(2,2) not null,
    sensor_id integer not null,
    FOREIGN KEY(sensor_id) REFERENCES sensors(id) ON DELETE CASCADE
);
`

func Connect(dbname string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", fmt.Sprintf("./%s.sqlite3", dbname))
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
