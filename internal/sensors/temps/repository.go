package temps

import (
	"fmt"
	"github.com/pavel-one/SimpleServerMonitor/internal/base"
	"github.com/pavel-one/SimpleServerMonitor/internal/charts"
	"github.com/pavel-one/SimpleServerMonitor/internal/db"
	"time"
)

var schema = `
CREATE TABLE IF NOT EXISTS Temps
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

func (r *Repository) Save(stat *Stat) error {
	q := `INSERT INTO Temps (key, temp) VALUES (:key, :temp)`

	_, err := r.DB.NamedExec(q, stat)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetBySeconds() ([]charts.Data, error) {
	var models []Model

	q := `
SELECT key, temp, time
FROM Temps
WHERE time BETWEEN datetime('now', '-1 minute') AND datetime('now')
ORDER BY time DESC
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return modelsToChart(models), nil
}

func (r *Repository) GetByMinutes() ([]charts.Data, error) {
	var models []Model

	q := `
SELECT key, temp, time
FROM Temps
WHERE time BETWEEN datetime('now', '-1 hour') AND datetime('now')
ORDER BY time DESC
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return modelsToChart(models), nil
}

func (r *Repository) GetByHours() ([]charts.Data, error) {
	var models []Model

	q := `
SELECT key, temp, time
FROM Temps
WHERE time BETWEEN datetime('now', '-1 day') AND datetime('now')
ORDER BY time DESC
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return modelsToChart(models), nil
}

func (r *Repository) GetByDays() ([]charts.Data, error) {
	var models []Model

	q := `
SELECT key, avg(temp) as temp, time
FROM Temps
WHERE time BETWEEN datetime('now', '-1 month') AND datetime('now')
GROUP BY strftime('%Y-%m-%d %H:00:00', time), key
ORDER BY time DESC
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return modelsToChart(models), nil
}

func (r *Repository) GetByMonth() ([]charts.Data, error) {
	var models []Model

	q := `
SELECT key, avg(temp) as temp, time
FROM Temps
WHERE time BETWEEN datetime('now', '-1 year') AND datetime('now')
GROUP BY strftime('%Y-%m-%d', time), key
ORDER BY time DESC
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return modelsToChart(models), nil
}
