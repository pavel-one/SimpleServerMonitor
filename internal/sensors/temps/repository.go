package temps

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/SimpleServerMonitor/internal/db"
)

type Repository struct {
	DB *sqlx.DB
}

func NewRepository() (*Repository, error) {
	connection, err := db.DefaultConnection()
	if err != nil {
		return nil, err
	}

	return &Repository{DB: connection}, nil
}
