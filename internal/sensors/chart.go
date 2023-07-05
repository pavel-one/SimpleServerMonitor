package sensors

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type sqlModel struct {
	Time     string  `db:"time"`
	Temp     float32 `db:"temp"`
	Name     string  `db:"name"`
	SensorId uint    `db:"sensor_id"`
}

type dataset struct {
	Label    string    `json:"label"`
	SensorId uint      `json:"sensor_id"`
	Data     []float32 `json:"data"`
}

type Chart struct {
	Labels   []string   `json:"labels"`
	Datasets []*dataset `json:"datasets"`
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
SELECT strftime('%Y-%m-%d %H:%M:%S', created_at) AS time,
       round(AVG(temp), 2) AS temp, 
       name,
       sensor_id
FROM sensors_data
INNER JOIN sensors s on s.id = sensors_data.sensor_id
WHERE created_at >= datetime('now', '-1 minute')
GROUP BY time, sensor_id
ORDER BY time;
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return r.ModelsToChart(models, "15:04:05"), nil
}

func (r *ChartSensorsRepository) ByMinutes() (*Chart, error) {
	var models []*sqlModel

	q := `
SELECT strftime('%Y-%m-%d %H:%M:00', created_at) AS time,
       round(AVG(temp), 2) AS temp, 
       name,
       sensor_id
FROM sensors_data
INNER JOIN sensors s on s.id = sensors_data.sensor_id
WHERE created_at >= datetime('now', '-1 hour')
GROUP BY time, sensor_id
ORDER BY time;
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return r.ModelsToChart(models, "15:04"), nil
}

func (r *ChartSensorsRepository) ByHours() (*Chart, error) {
	var models []*sqlModel

	q := `
SELECT strftime('%Y-%m-%d %H:00:00', created_at) AS time,
       round(AVG(temp), 2) AS temp, 
       name,
       sensor_id
FROM sensors_data
INNER JOIN sensors s on s.id = sensors_data.sensor_id
WHERE created_at >= datetime('now', '-24 hours')
GROUP BY time, sensor_id
ORDER BY time;
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return r.ModelsToChart(models, "2006-01-02 15:00"), nil
}

func (r *ChartSensorsRepository) ByDays() (*Chart, error) {
	var models []*sqlModel

	q := `
SELECT strftime('%Y-%m-%d 00:00:00', created_at) AS time,
       round(AVG(temp), 2) AS temp, 
       name,
       sensor_id
FROM sensors_data
INNER JOIN sensors s on s.id = sensors_data.sensor_id
WHERE created_at >= datetime('now', '-30 days')
GROUP BY time, sensor_id
ORDER BY time;
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return r.ModelsToChart(models, "2006-01-02"), nil
}

func (r *ChartSensorsRepository) ByMonth() (*Chart, error) {
	var models []*sqlModel

	q := `
SELECT strftime('%Y-%m-01 00:00:00', created_at) AS time,
       round(AVG(temp), 2) AS temp, 
       name,
       sensor_id
FROM sensors_data
INNER JOIN sensors s on s.id = sensors_data.sensor_id
WHERE created_at >= datetime('now', '-1 year')
GROUP BY time, sensor_id
ORDER BY time;
`

	if err := r.DB.Select(&models, q); err != nil {
		return nil, err
	}

	return r.ModelsToChart(models, "2006-01-02"), nil
}

func (r *ChartSensorsRepository) ModelsToChart(models []*sqlModel, timeLayout string) *Chart {
	mapper := make(map[uint]*dataset)
	mapperTimes := make(map[time.Time]bool)
	var labels []string
	var datasets []*dataset

	for _, m := range models {
		t, err := time.Parse("2006-01-02 15:04:05", m.Time)
		if err != nil {
			continue
		}

		v, ok := mapper[m.SensorId]
		if ok {
			v.Data = append(v.Data, m.Temp)
		} else {
			mapper[m.SensorId] = &dataset{
				Label:    m.Name,
				SensorId: m.SensorId,
				Data:     []float32{m.Temp},
			}
		}

		_, ok = mapperTimes[t]
		if !ok {
			mapperTimes[t] = true
		}
	}

	for t := range mapperTimes {
		labels = append(labels, t.Format(timeLayout))
	}

	for _, ds := range mapper {
		datasets = append(datasets, ds)
	}

	chart := &Chart{
		Labels:   labels,
		Datasets: datasets,
	}

	return chart
}
