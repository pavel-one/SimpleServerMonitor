package sensors

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type sqlModel struct {
	Time time.Time `db:"time"`
	Temp float32   `db:"temp"`
	Name string    `db:"name"`
}

type dataset struct {
	Label string    `json:"label"`
	Data  []float32 `json:"data"`
}

type Chart struct {
	Labels   []string  `json:"labels"`
	Datasets []dataset `json:"datasets"`
}

type ChartSensorsRepository struct {
	DB *sqlx.DB
}

func NewChartSensorsRepository(DB *sqlx.DB) *ChartSensorsRepository {
	return &ChartSensorsRepository{DB: DB}
}

func (r *ChartSensorsRepository) BySeconds() (*Chart, error) {
	var models []*sqlModel

	q := `
SELECT created_at AS time,
       AVG(temp) AS temp, name
FROM sensors_data
INNER JOIN sensors s on s.id = sensors_data.sensor_id
WHERE created_at >= datetime('now', '-1 minute')
GROUP BY time, sensor_id
ORDER BY time;
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return r.ModelsToChart(models), nil
}

func (r *ChartSensorsRepository) ByMinutes() {

}

func (r *ChartSensorsRepository) ByHours() {

}

func (r *ChartSensorsRepository) ByDays() {

}

func (r *ChartSensorsRepository) ByMonth() {

}

func (r *ChartSensorsRepository) ModelsToChart(models []*sqlModel) *Chart {
	return nil
}
