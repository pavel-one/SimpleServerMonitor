package charts

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Repository chart database
type Repository struct {
	DB *sqlx.DB
}

// NewRepository create new repository
func NewRepository(DB *sqlx.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) getQuery(timeFormat string, timeOffset string) string {
	q := `
SELECT strftime('%s', created_at) AS time,
       round(AVG(temp), 2)                       AS temp,
       c.name || ' [' || s.name || '] '          AS name,
       sensor_id
FROM sensors_data
         INNER JOIN sensors s on s.id = sensors_data.sensor_id
         INNER JOIN chips c on c.id = s.chip_id
WHERE created_at >= datetime('now', '%s')
GROUP BY time, sensor_id
ORDER BY time;
`

	return fmt.Sprintf(q, timeFormat, timeOffset)
}

// BySeconds getting chart for last 1 minute
func (r *Repository) BySeconds() (*Chart, error) {
	var models []*Model

	q := r.getQuery("%Y-%m-%d %H:%M:%S", "-1 minute")

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return mapToChart(models, "15:04:05"), nil
}

// ByMinutes getting chart for last 1 hour
func (r *Repository) ByMinutes() (*Chart, error) {
	var models []*Model

	q := r.getQuery("%Y-%m-%d %H:%M:00", "-1 hour")

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return mapToChart(models, "15:04"), nil
}

// ByHours getting chart for last 1 day
func (r *Repository) ByHours() (*Chart, error) {
	var models []*Model

	q := r.getQuery("%Y-%m-%d %H:00:00", "-1 day")

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return mapToChart(models, "2006-01-02 15:00"), nil
}

// ByDays getting chart for last 1 month
func (r *Repository) ByDays() (*Chart, error) {
	var models []*Model

	q := r.getQuery("%Y-%m-%d 00:00:00", "-1 month")

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return mapToChart(models, "2006-01-02"), nil
}

// ByMonth getting chart for last 1 year
func (r *Repository) ByMonth() (*Chart, error) {
	var models []*Model

	q := r.getQuery("%Y-%m-01 00:00:00", "-1 year")

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return mapToChart(models, "2006-01-02"), nil
}
