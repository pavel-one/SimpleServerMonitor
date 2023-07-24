package temps

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/SimpleServerMonitor/internal/db"
)

var schema = `
CREATE TABLE IF NOT EXISTS Temps
(
    id		VARCHAR NOT NULL PRIMARY KEY,
    temp 	INTEGER NOT NULL,
    time 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type StatRepository struct {
	DB *sqlx.DB
}

func NewStatRepository() (*StatRepository, error) {
	connection, err := db.DefaultConnection()
	if err != nil {
		return nil, err
	}

	if _, err := connection.Exec(schema); err != nil {
		return nil, fmt.Errorf("error create schema: %s", err)
	}

	return &StatRepository{DB: connection}, nil
}

func (r *StatRepository) Save(stat *Stat) error {

}
